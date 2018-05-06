package main

import (
	"github.com/my0sot1s/social/db"
	"github.com/my0sot1s/social/redis"
	"github.com/my0sot1s/social/utils"
	yaml "gopkg.in/yaml.v2"
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
