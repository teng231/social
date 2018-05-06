package main

import (
	"os"
	"strconv"
	"sync"

	"github.com/my0sot1s/social/api/services"
	"github.com/my0sot1s/social/core"
	"github.com/my0sot1s/social/db"
	"github.com/my0sot1s/social/mail"
	"github.com/my0sot1s/social/redis"
	"github.com/my0sot1s/social/utils"
)

func main() {
	c := loadConfig()
	if os.Getenv("ENV_MODE") == "production" {
		c.HOST = "lit-eyrie-97480.herokuapp.com"
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	mg := &db.DB{}
	go func() {
		mg.Config(c.DbHost, c.DbName, c.Username, c.Password)
		wg.Done()
	}()
	// register Redis cache
	rdCli := &redis.RedisCli{}
	go func() {
		err := rdCli.Config(c.RedisHost, c.RedisDB, c.RedisPass)
		if err != nil {
			utils.ErrLog(err)
		}
		wg.Done()
	}()

	mailCtrl := &mail.EmailMgr{}
	go func() {
		mailPort, e := strconv.Atoi(c.EmailPort)
		mailCtrl.Config(c.EmailHost, c.EmailOwn, c.EmailPw, mailPort)
		if e != nil {
			utils.ErrLog(e)
		}
		wg.Done()
	}()
	wg.Wait()
	// close when fn main down
	defer beforeDestroy(mg, rdCli)
	// create Core
	mainCore := &core.Social{}
	mainCore.Config(c.HOST, mg, rdCli, mailCtrl, "keys/id_rsa", "keys/id_rsa.pub")

	// create RESTful
	port := os.Getenv("PORT")
	if port == "" {
		port = c.PORT
	}
	RESTful := &api.GinConfig{}
	RESTful.Config(port, "", mainCore)
	RESTful.Run()

}
