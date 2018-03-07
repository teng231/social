package db

import "gopkg.in/mgo.v2/bson"

func (db *DB) GetPost(limit, page int) []*Post {
	collection := db.db.C(postCollection)
	var post []*Post
	query := collection.Find(bson.M{}).Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&post)
	if err != nil {
		panic(err)
	}
	return post
}

func (db *DB) GetFeed(limit, page int) []*Feed {
	collection := db.db.C(feedCollection)
	var feed []*Feed
	query := collection.Find(bson.M{}).Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&feed)
	if err != nil {
		panic(err)
	}
	return feed
}

func (db *DB) GetFollower(limit, page int, own string) []*Follower {
	collection := db.db.C(followerCollection)
	var follower []*Follower
	query := collection.Find(bson.M{"own": own}).Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&follower)
	if err != nil {
		panic(err)
	}
	return follower
}

func (db *DB) GetFollowing(limit, page int, follower string) []*Follower {
	collection := db.db.C(followerCollection)
	var following []*Follower
	query := collection.Find(bson.M{"follower": follower, "state": true})
	query = query.Skip(limit * (page - 1)).Limit(limit)
	err := query.All(&following)
	if err != nil {
		panic(err)
	}
	return following
}

func (db *DB) CountLike(postID string) int {
	collection := db.db.C(likeCollection)
	count, err := collection.Find(bson.M{"post_id": postID, "state": true}).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func (db *DB) GetAlbum(AlbumID string) *Album {
	collection := db.db.C(albumCollection)
	var album *Album
	err := collection.FindId(AlbumID).One(&album)
	if err != nil {
		panic(err)
	}
	return album
}
