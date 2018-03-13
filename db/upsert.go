package db

import (
	"fmt"
	"time"

	m "github.com/my0sot1s/social/mongo"

	"gopkg.in/mgo.v2/bson"
)

func (db *DB) CreatePost(p *m.Post) (error, *m.Post) {
	collection := db.Db.C(postCollection)
	p.ID = bson.NewObjectId()
	err := collection.Insert(&p)
	if err != nil {
		return err, nil
	}
	return nil, p
}

func (db *DB) CreateComment(c *m.Comment) (error, *m.Comment) {
	collection := db.Db.C(commentCollection)
	c.ID = bson.NewObjectId()
	err := collection.Insert(&c)
	if err != nil {
		return err, nil
	}
	return nil, c
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
		feedDone = append(feedDone, *feed)
	}
	bulk.Insert(feedDone...)
	_, bulkErr := bulk.Run()
	if bulkErr != nil {
		return bulkErr, nil
	}
	return nil, feedDone
}

func (db *DB) CreateUser(u *m.User) (error, *m.User) {
	collection := db.Db.C(userCollection)
	u.ID = bson.NewObjectId()
	u.AlbumName = fmt.Sprintf("@%s_%s", u.GetUserName(), u.GetID())
	err := collection.Insert(&u)
	if err != nil {
		return err, nil
	}
	return nil, u
}

func (db *DB) ModifyFollower(t *m.Follower) (error, *m.Follower) {
	collection := db.Db.C(followerCollection)
	t.ID = bson.NewObjectId()
	err := collection.Insert(&t)
	if err != nil {
		return err, nil
	}
	return nil, t
}

func (db *DB) CreateAlbum(a *m.Album) (error, *m.Album) {
	collection := db.Db.C(albumCollection)
	a.ID = bson.NewObjectId()
	err := collection.Insert(&a)
	if err != nil {
		return err, nil
	}
	return nil, a
}

func (db *DB) HitLikePost(postID, userID string) error {
	collection := db.Db.C(likeCollection)
	sellector := bson.M{"post_id": postID, "user_id": userID}
	update := bson.M{
		"created": time.Now(),
		"post_id": postID,
		"user_id": userID,
	}
	_, err := collection.Upsert(sellector, update)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UnlikePost(postID, userID string) error {
	collection := db.Db.C(likeCollection)
	sellector := bson.M{"post_id": postID, "user_id": userID}
	err := collection.Remove(sellector)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetUsersLikePost(userIDs []string) (error, []*m.User) {
	collection := db.Db.C(userCollection)
	var users []*m.User
	listQueriesID := make([]bson.M, 0)
	for _, p := range userIDs {
		if p == "" {
			continue
		}
		bsonID := bson.M{"_id": bson.ObjectIdHex(p)}
		listQueriesID = append(listQueriesID, bsonID)
	}
	err := collection.Find(bson.M{"$or": listQueriesID}).All(&users)
	if err != nil {
		return err, nil
	}
	return nil, users
}

// `owner` follow `uid`
func (db *DB) FollowUser(f *m.Follower) error {
	collection := db.Db.C(followerCollection)
	err := collection.Insert(f)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UnfollowUser(own, uid string) error {
	collection := db.Db.C(followerCollection)
	err := collection.Remove(bson.M{"own": uid, "follower": own})
	if err != nil {
		return err
	}
	return nil
}
