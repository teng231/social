package main

import (
	"io/ioutil"
	"social/api"
	"social/core"
	"social/db"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DbHost   string `yaml:"mgo_host" required:"true"`
	DbName   string `yaml:"mgo_database" required:"true"`
	Username string `yaml:"mgo_uname" required:"true"`
	Password string `yaml:"mgo_pword" required:"true"`
	PORT     int    `yaml:"gin_PORT" required:"true"`
}

func loadConfig() *Config {
	t := &Config{}
	yamlText, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlText, t)
	if err != nil {
		panic(err)
	}

	return t
}

func main() {
	c := loadConfig()
	mongo := &db.DB{}
	mongo.Config(c.DbHost, c.DbName, c.Username, c.Password)
	// close when fn main down
	defer mongo.Close()

	core := &core.Core{}
	core.Config(mongo)
	core.GetPostByUser(10, 1, "5a106155cb8eae85d819a78d")
	// core.GetFeedByUser(10, 1, "5a106166cb8eae85d819a78e")
	RESTful := &api.GinConFig{}
	RESTful.Config(c.PORT, "")
	RESTful.Run()
}
