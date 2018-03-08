package core

import (
	"errors"
	"fmt"
	"log"
	m "social/mongo"

	"golang.org/x/crypto/bcrypt"
)

type IAuth interface {
	GetUserByUname(username string) *m.User
	CreateUser(u *m.User) (error, *m.User)
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

func (c *Core) Login(username, password string) (error, interface{}) {
	err, user := c.Db.GetUserByUname(username)
	if err != nil {
		return err, nil
	}
	if user == nil {
		return errors.New("user khong ton tai"), nil
	}
	if !comparePassword(password, user.GetPassword()) {
		return errors.New("password false"), nil
	}
	return errors.New("login success"), nil
}

func (c *Core) Register(user *m.User) (error, *m.User) {
	username := user.GetUserName()
	if username == "" || user.GetEmail() == "" || user.GetPassword() == "" {
		return errors.New("No username or email or password"), nil
	}
	err, _user := c.Db.GetUserByUname(username)

	if _user != nil {
		return errors.New("username is exist"), nil
	}

	err, email := c.Db.GetUserByEmail(user.GetEmail())

	if email != nil {
		return errors.New("email is exist"), nil
	}

	if err != nil {
		return err, nil
	}

	if len(user.GetPassword()) < 6 {
		return errors.New("password short"), nil
	}
	return nil, nil
}
