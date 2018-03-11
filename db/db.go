package db

import (
	"time"

	"github.com/my0sot1s/social/utils"

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
	albumCollection    = "social_albums"
	commentCollection  = "social_comments"
	followerCollection = "social_follower"
	likeCollection     = "social_like"
	tokenCollection    = "AccessToken"
	feedCollection     = "social_feed"
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
		utils.ErrLog(err)
		return
	}
	connect.Session = mongoSession
	mongoSession.SetMode(mgo.Monotonic, true)
	connect.Db = mongoSession.DB(Database)
	utils.Log("ಠ‿ಠ mongodb connected ಠ‿ಠ")
}

func (connect *DB) Close() {
	connect.Session.Close()
}
