package db

import (
	"social/utils"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// DB
type DB struct {
	Session *mgo.Session
	Db      *mgo.Database
}

const (
	postCollection     = "social_post"
	userCollection     = "social_user"
	albumCollection    = "social_album"
	commentCollection  = "social_collection"
	followerCollection = "social_follower"
	likeCollection     = "social_like"
	tokenCollection    = "AccessToken"
	feedCollection     = "social_timeline"
)

func (connect *DB) Config(Host, Database, Username, Password string) {
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{Host},
		Timeout:  60 * time.Second,
		Database: Database,
		Username: Username,
		Password: Password,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		panic(err)
	}
	connect.Session = mongoSession
	mongoSession.SetMode(mgo.Monotonic, true)
	if err = mongoSession.DB(Database).Login(Username, Password); err != nil {
		panic(err)
	}
	connect.Db = mongoSession.DB(Database)
	utils.Log("** Config success")
}

func (connect *DB) Close() {
	connect.Session.Close()
}
