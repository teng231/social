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

	"gopkg.in/yaml.v2"
)

type Config struct {
	GO_MODE   string `yaml:"GO_MODE" required:"true"`
	HOST      string `yaml:"HOST" required:"true"`
	DbHost    string `yaml:"mgo_Host" required:"true"`
	DbName    string `yaml:"mgo_Database" required:"true"`
	Username  string `yaml:"mgo_Username" required:"true"`
	Password  string `yaml:"mgo_Password" required:"true"`
	PORT      string `yaml:"gin_PORT" required:"true"`
	RedisHost string `yaml:"redis_Host" required:"true"`
	RedisDB   string `yaml:"redis_Db" required:"true"`
	RedisPass string `yaml:"redis_Password" required:"true"`
	EmailHost string `yaml:"email_Host" required:"true"`
	EmailOwn  string `yaml:"email" required:"true"`
	EmailPw   string `yaml:"email_P" required:"true"`
	EmailPort string `yaml:"email_Port" required:"true"`
}

func loadConfig() *Config {
	t := &Config{}
	yamlText, err := utils.ReadFileRoot("config.yaml")
	err = yaml.Unmarshal(yamlText, t)
	if err != nil {
		utils.ErrLog(err)
		return nil
	}

	return t
}
func beforeDestroy(mg *db.DB, rc *redis.RedisCli) {
	if r := recover(); r != nil {
		utils.Log(r)
		return
	}
	utils.Log("app exited")
	rc.Close()
	mg.Close()
}
func main() {
	c := loadConfig()
	if os.Getenv("ENV_MODE") == "production" {
		c.HOST = "serene-headland-81432.herokuapp.com"
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
	mainCore := &core.Core{}
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
