package db

import (
	m "social/mongo"
	"social/utils"

	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetUserByUname(username string) *m.User {
	collection := db.Db.C(postCollection)
	user := &m.User{}
	err := collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func (db *DB) GetUserByEmail(email string) *m.User {
	collection := db.Db.C(postCollection)
	user := &m.User{}
	err := collection.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func (db *DB) GetPost(limit, page int, userId string) []*m.Post {
	utils.Log(userId)
	collection := db.Db.C(postCollection)
	var post []*m.Post
	query := collection.Find(bson.M{"user_id": userId})
	err := query.Skip(limit * (page - 1)).Limit(limit).All(&post)
	if err != nil {
		panic(err)
	}
	return post
}

func (db *DB) GetPostById(postID string) *m.Post {
	utils.Log(postID)
	collection := db.Db.C(postCollection)
	var post *m.Post
	err := collection.FindId(postID).One(&post)
	if err != nil {
		panic(err)
	}
	return post
}

func (db *DB) GetFeed(limit, page int, userId string) []*m.Feed {
	collection := db.Db.C(feedCollection)
	var feed []*m.Feed
	query := collection.Find(bson.M{"comsumer_id": userId})
	query = query.Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&feed)
	if err != nil {
		panic(err)
	}
	return feed
}

func (db *DB) GetFollower(limit, page int, own string) []*m.Follower {
	collection := db.Db.C(followerCollection)
	var follower []*m.Follower
	query := collection.Find(bson.M{"own": own})
	query = query.Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&follower)
	if err != nil {
		panic(err)
	}
	return follower
}

func (db *DB) GetFollowing(limit, page int, follower string) []*m.Follower {
	collection := db.Db.C(followerCollection)
	var following []*m.Follower
	query := collection.Find(bson.M{"follower": follower, "state": true})
	query = query.Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&following)
	if err != nil {
		panic(err)
	}
	return following
}

func (db *DB) CountLike(postID string) int {
	collection := db.Db.C(likeCollection)
	count, err := collection.Find(bson.M{"post_id": postID, "state": true}).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func (db *DB) GetAlbum(AlbumID string) *m.Album {
	collection := db.Db.C(albumCollection)
	var album *m.Album
	err := collection.FindId(AlbumID).One(&album)
	if err != nil {
		panic(err)
	}
	return album
}
