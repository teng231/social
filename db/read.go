package db

import (
	"fmt"

	"github.com/my0sot1s/social/utils"

	m "github.com/my0sot1s/social/mongo"

	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetUserByUname(username string) (error, *m.User) {
	collection := db.Db.C(postCollection)
	user := &m.User{}
	err := collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (db *DB) GetUserByEmail(email string) (error, *m.User) {
	collection := db.Db.C(postCollection)
	user := &m.User{}
	err := collection.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (db *DB) GetPost(limit, page int, userId string) (error, []*m.Post) {
	utils.Log(userId)
	collection := db.Db.C(postCollection)
	var post []*m.Post
	query := collection.Find(bson.M{"user_id": userId})
	err := query.Skip(limit * (page - 1)).Limit(limit).All(&post)
	if err != nil {
		return err, nil
	}
	return nil, post
}

func (db *DB) GetPostById(postID string) (error, *m.Post) {
	utils.Log(postID)
	collection := db.Db.C(postCollection)
	var post *m.Post
	err := collection.FindId(bson.ObjectIdHex(postID)).One(&post)
	if err != nil {
		return err, nil
	}
	return nil, post
}

func (db *DB) GetFeed(limit, page int, userId string) (error, []*m.Feed) {
	collection := db.Db.C(feedCollection)
	var feeds []*m.Feed
	query := collection.Find(bson.M{"comsumer_id": userId})
	query = query.Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&feeds)
	if err != nil {
		return err, nil
	}
	return nil, feeds
}

func (db *DB) GetFollower(limit, page int, own string) (error, []*m.Follower) {
	collection := db.Db.C(followerCollection)
	var follower []*m.Follower
	query := collection.Find(bson.M{"own": own})
	query = query.Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&follower)
	if err != nil {
		return err, nil
	}
	return nil, follower
}

func (db *DB) GetFollowing(limit, page int, follower string) (error, []*m.Follower) {
	collection := db.Db.C(followerCollection)
	var following []*m.Follower
	query := collection.Find(bson.M{"follower": follower, "state": true})
	query = query.Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&following)
	if err != nil {
		return err, nil
	}
	return nil, following
}

func (db *DB) CountLike(postID string) (error, int) {
	collection := db.Db.C(likeCollection)
	count, err := collection.Find(bson.M{"post_id": postID, "state": true}).Count()
	if err != nil {
		return err, 0
	}
	return nil, count
}

func (db *DB) GetAlbum(AlbumID string) (error, *m.Album) {
	collection := db.Db.C(albumCollection)
	var album *m.Album
	err := collection.FindId(bson.ObjectIdHex(AlbumID)).One(&album)
	if err != nil {
		return err, nil
	}
	return nil, album
}

func (db *DB) GetComments(limit, page int, postID string) (error, []*m.Comment) {
	collection := db.Db.C(commentCollection)
	var comments []*m.Comment
	err := collection.Find(bson.M{"post_id": postID}).Skip(limit * (page - 1)).Limit(limit).All(&comments)
	if err != nil {
		return err, nil
	}
	fmt.Printf("%v", len(comments))
	return nil, comments
}
