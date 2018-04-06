package core

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
	"golang.org/x/crypto/bcrypt"
)

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

func (c *Social) Login(username, password string) (error, *m.User, string) {
	err, user := c.Db.GetUserByUname(username)
	if err != nil {
		utils.ErrLog(err)
		return err, nil, ""
	}
	if user == nil {
		utils.ErrLog(err)
		return errors.New("user khong ton tai"), nil, ""
	}
	if user.GetState() != "activated" {
		utils.ErrLog(errors.New("account not actived"))
		return errors.New("account not actived"), nil, ""
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

func (c *Social) Logout(token string) error {
	var keyList = []string{token}
	_, err := c.rd.DelKey(keyList)
	utils.ErrLog(err)
	return err
}

func (c *Social) Register(user *m.User) (error, *m.User) {
	username := user.GetUserName()

	if username == "" || user.GetEmail() == "" || user.GetPassword() == "" {
		utils.ErrLog(errors.New("No username or email or password"))
		return errors.New("No username or email or password"), nil
	}
	err, _user := c.Db.GetUserByUname(username)

	if _user != nil {
		utils.ErrLog(errors.New("username is exist"))
		return errors.New("username is exist"), nil
	}

	utils.Log(user.GetEmail())
	err, email := c.Db.GetUserByEmail(user.GetEmail())

	if email != nil {

		utils.ErrLog(errors.New("email is exist"))
		return errors.New("email is exist"), nil
	}

	if len(user.GetPassword()) < 6 {

		utils.ErrLog(errors.New("password short"))
		return errors.New("password short"), nil
	}
	user.SetPassword(generateHashPassword(user.GetPassword()))
	user.Created = time.Now()
	user.State = "pendding"
	e := c.Db.CreateUser(user)

	if e != nil {
		utils.ErrLog(e)
		return err, nil
	}
	_, tokenString := c.token.GenerateToken(user.GetID())
	err = c.rd.SetValue(tokenString, user.GetID(), defaultExpireOffset*48) // a day
	utils.Log(user.GetID())
	utils.Log(tokenString)
	// send email
	// read Email
	data, err := utils.ReadFileRoot("./mail/confirm.htm")
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	link := fmt.Sprintf("http://%s/confirm/%s/%s", c.HOST, user.GetID(), tokenString)
	emailContent := strings.Replace(string(data), "##link##", link, 1)
	c.mailAd.SendMail(c.mailAd.Username, emailContent, "Confirm create Account", user.GetEmail())
	return e, user
}

func (c *Social) ActivedAccount(uid string) error {
	err := c.Db.UpdateStateUser(uid, "activated")
	if err != nil {
		utils.ErrLog(err)
		return err
	}
	return nil
}

func (c *Social) ChangePassword(username, oldPass, newPass string) error {
	err, user := c.Db.GetUserByUname(username)
	if err != nil {
		utils.ErrLog(err)
		return err
	}
	if generateHashPassword(oldPass) == user.GetPassword() {
		err2 := c.Db.UpdateUserPassword(user.GetID(), generateHashPassword(newPass))
		if err2 != nil {
			utils.ErrLog(err2)
			return err2
		}
	}
	return nil
}

func (c *Social) getUserByIDs(uIDs []string) (error, []*m.User) {
	err, users := c.Db.GetUserByIds(uIDs)
	if err != nil {
		utils.ErrLog(err)
		return nil, nil
	}
	for _, v := range users {
		v.SetPassword("")
	}
	return err, users
}
func (c *Social) CheckKeyToken(token, uidOwn string) error {
	uid, err := c.rd.GetValue(token)
	if err != nil {
		utils.ErrLog(err)
		return err
	}
	if strings.Trim(uid, " ") != strings.Trim(uidOwn, " ") {
		return errors.New("uid not equal")
	}
	_, err2 := c.rd.DelKey([]string{token})
	if err2 != nil {
		utils.ErrLog(err2)
		return err2
	}
	utils.Log(uidOwn)
	err3 := c.Db.UpdateStateUser(uidOwn, "activated")
	if err3 != nil {
		utils.ErrLog(err3)
		return err3
	}
	return nil
}
