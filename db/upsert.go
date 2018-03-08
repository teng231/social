package db

import (
	m "social/mongo"

	"gopkg.in/mgo.v2/bson"
)

func (db *DB) CreatePost(p *m.Post) *m.Post {
	collection := db.Db.C(postCollection)
	p.ID = bson.NewObjectId()
	err := collection.Insert(&p)
	if err != nil {
		panic(err)
	}
	return p
}
func (db *DB) CreateUser(u *m.User) (error, *m.User) {
	collection := db.Db.C(userCollection)
	u.ID = bson.NewObjectId()
	err := collection.Insert(&u)
	if err != nil {
		panic(err)
	}
	return nil, u
}

func (db *DB) CreateComment(c *m.Comment) *m.Comment {
	collection := db.Db.C(commentCollection)
	c.ID = bson.NewObjectId()
	err := collection.Insert(&c)
	if err != nil {
		panic(err)
	}
	return c
}

func (db *DB) CreateFeed(f *m.Feed) *m.Feed {
	collection := db.Db.C(feedCollection)
	f.ID = bson.NewObjectId()
	err := collection.Insert(&f)
	if err != nil {
		panic(err)
	}
	return f
}

func (db *DB) CreateFeeds(feeds []*m.Feed) []interface{} {
	bulk := db.Db.C(feedCollection).Bulk()
	feedDone := make([]interface{}, 0)
	for _, feed := range feeds {
		feed.ID = bson.NewObjectId()
		feedDone = append(feedDone, *feed)
	}
	bulk.Insert(feedDone...)
	_, bulkErr := bulk.Run()
	if bulkErr != nil {
		panic(bulkErr)
	}
	return feedDone
}

func (db *DB) CreateUser(u *m.User) *m.User {
	collection := db.Db.C(userCollection)
	u.ID = bson.NewObjectId()
	err := collection.Insert(&u)
	if err != nil {
		panic(err)
	}
	return u
}

func (db *DB) CreateToken(t *m.AccessToken) *m.AccessToken {
	collection := db.Db.C(tokenCollection)
	t.ID = bson.NewObjectId()
	err := collection.Insert(&t)
	if err != nil {
		panic(err)
	}
	return t
}

// func (db *DB) ModifyLike(t *Like, postID, userID string) *Like {
// 	collection := db.db.C(tokenCollection)
// 	t.ID = bson.NewObjectId()
// 	updator := bson.M{"$set", t}
// 	err := collection.Upsert(bson.M{"post_id": postID, "user_id": userID}, updator)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return t
// }

func (db *DB) ModifyFollower(t *m.Follower) *m.Follower {
	collection := db.Db.C(tokenCollection)
	t.ID = bson.NewObjectId()
	err := collection.Insert(&t)
	if err != nil {
		panic(err)
	}
	return t
}

func (db *DB) CreateAlbum(a *m.Album) *m.Album {
	collection := db.Db.C(albumCollection)
	a.ID = bson.NewObjectId()
	err := collection.Insert(&a)
	if err != nil {
		panic(err)
	}
	return a
}
