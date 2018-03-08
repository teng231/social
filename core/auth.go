package core

import (
	"fmt"
	"log"
	m "social/mongo"

	"golang.org/x/crypto/bcrypt"
)

type IAuth interface {
	GetUserByUname(username string) *m.User
	CreateUser(u *User) error,*m.User
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

func (c *Core) Login(username, password string) interface{} {
	user := c.Db.GetUserByUname(username)
	if user == nil {
		return "user khong ton tai"
	}
	if !comparePassword(password, user.GetPassword()) {
		return "password false"
	}
	return "login success"
}

func (c *Core) Register(user *m.User) (error,*User) {
	username := user.GetUserName()
	if username == "" || user.GetEmail()== "" || user.GetPassword() == ""{
		return errors.New("No username or email or password"),nil
	}
	_user := c.Db.GetUserByUname(username)

	if _user != nil {
		return errors.New("username is exist"),nil		
	}

	_user := c.Db.GetUserByEmail(user.GetEmail())
	if _user != nil {
		return errors.New("email is exist"),nil		
	}
	
	if len(user.GetPassword()) < 6 {
		return errors.New("password short"),nil		 
	}

	// ok xu ly
	// user.SetPassword()
}
