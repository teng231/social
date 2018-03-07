package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// DB
type DB struct {
	session *mgo.Session
	db      *mgo.Database
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

func (db *DB) Config(Host, Database, Username, Password string) {
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
	db.session = mongoSession
	mongoSession.SetMode(mgo.Monotonic, true)
	db.db = mongoSession.DB(Database)

}
