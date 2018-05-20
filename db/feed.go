package db

import (
	"time"

	m "github.com/my0sot1s/social/mirrors"

	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetMigrateFeed(limit int) (error, []*m.Feed) {
	feeds := make([]*m.Feed, 0)
	err, mf := db.ReadByIdCondition(feedCollection, "", limit, bson.M{})
	if err != nil {
		return err, nil
	}
	for _, v := range mf {
		f := &m.Feed{}
		f.ToFeed(v)
		feeds = append(feeds, f)
	}
	return nil, feeds
}

func (db *DB) GetFeed(limit int, anchor, userId string) (error, []*m.Feed) {
	feeds := make([]*m.Feed, 0)
	// err, mf := db.ReadByIdCondition(feedCollection, anchor, limit, bson.M{"comsumer_id": userId})
	err, mf := db.ReadByIdOtherCondition("post_id", feedCollection, anchor, limit, bson.M{"comsumer_id": userId})
	if err != nil {
		return err, nil
	}
	for _, v := range mf {
		f := &m.Feed{}
		f.ToFeed(v)
		feeds = append(feeds, f)
	}
	return nil, feeds
}

func (db *DB) CreateFeed(f *m.Feed) error {
	collection := db.Db.C(feedCollection)
	f.ID = bson.NewObjectId()
	err := collection.Insert(&f)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateFeeds(feeds []*m.Feed) (error, []interface{}) {
	bulk := db.Db.C(feedCollection).Bulk()
	feedDone := make([]interface{}, 0)
	for _, feed := range feeds {
		feed.ID = bson.NewObjectId()
		feed.Created = time.Now()
		feedDone = append(feedDone, *feed)
	}
	bulk.Insert(feedDone...)
	_, bulkErr := bulk.Run()
	if bulkErr != nil {
		return bulkErr, nil
	}
	return nil, feedDone
}

func (db *DB) UpsertFeed(id string, f *m.Feed) error {
	collection := db.Db.C(feedCollection)
	collection.RemoveId(bson.ObjectIdHex(id))
	err := collection.Insert(*f)
	return err
}

func (db *DB) DeleteFeed(uid, owner string) error {
	collection := db.Db.C(feedCollection)
	_, err := collection.RemoveAll(bson.M{"author": owner, "comsumer_id": uid})
	return err
}
