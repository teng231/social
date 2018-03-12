package core

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
	"golang.org/x/crypto/bcrypt"
)

type IAuth interface {
	GetUserByUname(username string) *m.User
	CreateUser(u *m.User) (error, *m.User)
}

type JWTAuthentication struct {
	privateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	tokenDuration int
	expireOffset  int
}

const (
	defaultTokenDuration = 72
	defaultExpireOffset  = 30 * 60 * 60 // 30 min
)

func (jwta *JWTAuthentication) Config(privateKeyPath, PublicKeyPath string) {
	jwta.getPrivateKey(privateKeyPath)
	jwta.getPublicKey(PublicKeyPath)
	jwta.tokenDuration = defaultTokenDuration
	jwta.expireOffset = defaultExpireOffset
	utils.Log("  ಠ‿ಠ JWTAuthentication ready ಠ‿ಠ")
}

func (jwta *JWTAuthentication) GenerateToken(userUUID string) (error, string) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(jwta.tokenDuration)).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}
	tokenString, err := token.SignedString(jwta.privateKey)
	if err != nil {
		utils.ErrLog(err)
		return err, ""
	}
	return nil, tokenString
}

func (jwta *JWTAuthentication) GetTokenRemaining(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds()) + jwta.expireOffset
		}
	}
	return jwta.expireOffset
}

func (jwta *JWTAuthentication) readFileSSl(sslKeyPath string) []byte {
	pembytes, err := ioutil.ReadFile(sslKeyPath)
	utils.ErrLog(err)
	return pembytes
}

func (jwta *JWTAuthentication) getPublicKey(publicKeyPath string) {
	var err error
	file := jwta.readFileSSl(publicKeyPath)
	jwta.PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(file)
	if err != nil {
		utils.ErrLog(err)
		// return nil
	}
}

func (jwta *JWTAuthentication) getPrivateKey(privateKeyPath string) {
	var err error
	file := jwta.readFileSSl(privateKeyPath)
	jwta.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(file)
	if err != nil {
		utils.ErrLog(err)
	}
}

func generateHashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
		return ""
	}
	fmt.Println("Hash to store:", string(hash))
	return string(hash)
}

func comparePassword(userPassword, hashFromDatabase string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashFromDatabase), []byte(userPassword)); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
		return false
	}
	return true
}

func (c *Core) Login(username, password string) (error, *m.User, string) {
	err, user := c.Db.GetUserByUname(username)
	if err != nil {
		utils.ErrLog(err)
		return err, nil, ""
	}
	if user == nil {
		utils.ErrLog(err)
		return errors.New("user khong ton tai"), nil, ""
	}
	if !comparePassword(password, user.GetPassword()) {
		utils.ErrLog(err)
		return errors.New("password false"), nil, ""
	}
	_, tokenString := c.token.GenerateToken(user.GetID())
	user.SetPassword("")
	// insert to redis
	err = c.rd.SetValue(tokenString, user.GetID(), defaultExpireOffset)
	utils.ErrLog(err)
	return nil, user, tokenString
}

func (c *Core) Register(user *m.User) (error, *m.User, string) {
	username := user.GetUserName()

	if username == "" || user.GetEmail() == "" || user.GetPassword() == "" {
		utils.ErrLog(errors.New("No username or email or password"))
		return errors.New("No username or email or password"), nil, ""
	}
	err, _user := c.Db.GetUserByUname(username)

	if _user != nil {
		utils.ErrLog(errors.New("username is exist"))
		return errors.New("username is exist"), nil, ""
	}

	utils.Log(user.GetEmail())
	err, email := c.Db.GetUserByEmail(user.GetEmail())

	if email != nil {

		utils.ErrLog(errors.New("email is exist"))
		return errors.New("email is exist"), nil, ""
	}

	if len(user.GetPassword()) < 6 {

		utils.ErrLog(errors.New("password short"))
		return errors.New("password short"), nil, ""
	}
	user.SetPassword(generateHashPassword(user.GetPassword()))
	user.Created = time.Now()
	user.State = "actived"
	e, newUser := c.Db.CreateUser(user)

	if e != nil {
		utils.ErrLog(e)
		return err, nil, ""
	}
	_, tokenString := c.token.GenerateToken(newUser.GetID())
	newUser.SetPassword("")
	err = c.rd.SetValue(tokenString, user.GetID(), defaultExpireOffset)
	utils.Log(user.GetID())
	return e, newUser, tokenString
}

func (c *Core) CheckTokenExpired(tokenString string) string {
	uid, err := c.rd.GetValue(tokenString)

	if err != nil {
		utils.ErrLog(err)
		return ""
	}
	return uid
}

func (c *Core) ChangePassword(oldPass, newPass string) bool {
	return true
}

func (c *Core) getUserByIDs(uIDs []string) (error, []*m.User) {
	err, users := c.Db.GetUserOwns(uIDs)
	if err != nil {
		utils.ErrLog(err)
		return nil, nil
	}
	for _, v := range users {
		v.SetPassword("")
	}
	return err, users
}
