package db

import (
	"time"

	"github.com/my0sot1s/social/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	feedCollection     = "social_feed"
	saveCollection     = "social_saved"
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

// request: `condition`, `anchor`

func (db *DB) ReadByIdCondition(cName, anchor string, limit int, conditions map[string]interface{}) (error, []map[string]interface{}) {
	collection := db.Db.C(cName)
	result := make([]map[string]interface{}, 0)
	var query map[string]interface{}
	if anchor != "" {
		if limit < 0 {
			//  lt
			query = bson.M{"_id": bson.M{"$lt": bson.ObjectIdHex(anchor)}}
			utils.Log("run less than")
		} else {
			//  gt
			query = bson.M{"_id": bson.M{"$gt": bson.ObjectIdHex(anchor)}}
		}
	} else {
		query = bson.M{}
	}
	if conditions != nil {
		for k, c := range conditions {
			query[k] = c
		}
	}
	var err error

	q := collection.Find(query)
	if limit < 0 {
		err = q.Limit(-limit).Sort("-$natural").All(&result)
		utils.Log("run with limit < 0")
	} else {
		err = q.Limit(limit).Sort("-$natural").All(&result)
	}

	if err != nil {
		return err, nil
	}
	return nil, result
}

func (db *DB) ReadOneBasic(cName string, conditionQuery interface{}) (error, map[string]interface{}) {
	collection := db.Db.C(cName)
	var result map[string]interface{}
	err := collection.Find(conditionQuery).One(&result)
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (db *DB) ReadById(cName, anyId string) (error, map[string]interface{}) {
	collection := db.Db.C(cName)
	var result map[string]interface{}
	err := collection.FindId(bson.ObjectIdHex(anyId)).One(&result)
	if err != nil {
		return err, nil
	}
	return nil, result
}
