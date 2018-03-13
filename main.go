package main

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/my0sot1s/social/api"
	"github.com/my0sot1s/social/core"
	"github.com/my0sot1s/social/db"
	"github.com/my0sot1s/social/redis"
	"github.com/my0sot1s/social/utils"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DbHost    string `yaml:"mgo_Host" required:"true"`
	DbName    string `yaml:"mgo_Database" required:"true"`
	Username  string `yaml:"mgo_Username" required:"true"`
	Password  string `yaml:"mgo_Password" required:"true"`
	PORT      string `yaml:"gin_PORT" required:"true"`
	RedisHost string `yaml:"redis_Host" required:"true"`
	RedisDB   string `yaml:"redis_Db" required:"true"`
	RedisPass string `yaml:"redis_Password" required:"true"`
}

func loadConfig() *Config {
	t := &Config{}
	yamlText, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		utils.ErrLog(err)
		return nil
	}
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
	// register Db
	wg := sync.WaitGroup{}
	wg.Add(2)
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
	wg.Wait()
	// close when fn main down
	defer beforeDestroy(mg, rdCli)
	// create Core
	core := &core.Core{}
	core.Config(mg, rdCli, "keys/id_rsa", "keys/id_rsa.pub")
	// create RESTful
	port := os.Getenv("PORT")
	if port == "" {
		port = c.PORT
	}
	RESTful := &api.GinConfig{}
	RESTful.Config(port, "", core)
	RESTful.Run()
}
