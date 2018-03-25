package db

import (
	"time"

	m "github.com/my0sot1s/social/mongo"

	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetFeed(limit int, anchor, userId string) (error, []*m.Feed) {
	feeds := make([]*m.Feed, 0)
	err, mf := db.ReadByIdCondition(feedCollection, anchor, limit, bson.M{"comsumer_id": userId})
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

func (db *DB) CreateFeed(f *m.Feed) (error, *m.Feed) {
	collection := db.Db.C(feedCollection)
	f.ID = bson.NewObjectId()
	err := collection.Insert(&f)
	if err != nil {
		return err, nil
	}
	return nil, f
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
