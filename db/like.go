package db

import (
	"time"

	m "github.com/my0sot1s/social/mirrors"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) CountLike(postID string) (error, int) {
	collection := db.Db.C(likeCollection)
	count, err := collection.Find(bson.M{"post_id": postID}).Count()
	if err != nil {
		return err, 0
	}
	return nil, count
}

func (db *DB) GetLikes(postID string) (error, []*m.Like) {
	collection := db.Db.C(likeCollection)
	var likes []*m.Like
	err := collection.Find(bson.M{"post_id": postID}).All(&likes)
	if err != nil {
		return err, nil
	}
	return nil, likes
}
func (db *DB) IsUserLikePost(pid, uid string) (error, bool) {
	collection := db.Db.C(likeCollection)
	count, err := collection.Find(bson.M{"post_id": pid, "user_id": uid}).Count()
	if err != nil {
		return err, true
	}
	isExist := false
	if count > 0 {
		isExist = true
	}
	return nil, isExist
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

