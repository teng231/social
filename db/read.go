package db

import (
	m "github.com/my0sot1s/social/mongo"

	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetUserByUname(username string) (error, *m.User) {
	collection := db.Db.C(userCollection)
	user := &m.User{}
	err := collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (db *DB) GetUserByEmail(email string) (error, *m.User) {
	collection := db.Db.C(userCollection)
	user := &m.User{}
	err := collection.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (db *DB) GetPost(limit, page int, userId string) (error, []*m.Post) {
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

func (db *DB) GetFollower(own string) (error, []*m.Follower) {
	collection := db.Db.C(followerCollection)
	var follower []*m.Follower
	err := collection.Find(bson.M{"own": own}).All(&follower)
	if err != nil {
		return err, nil
	}
	return nil, follower
}

func (db *DB) GetFollowing(follower string) (error, []*m.Follower) {
	collection := db.Db.C(followerCollection)
	var following []*m.Follower
	err := collection.Find(bson.M{"follower": follower}).All(&following)
	if err != nil {
		return err, nil
	}
	return nil, following
}

func (db *DB) CountLike(postID string) (error, int) {
	collection := db.Db.C(likeCollection)
	count, err := collection.Find(bson.M{"post_id": postID}).Count()
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

func (db *DB) GetAlbumByAuthor(limit, page int, userId string) (error, []*m.Album) {
	collection := db.Db.C(albumCollection)
	var albums []*m.Album
	err := collection.Find(bson.M{"author": userId}).Skip(limit * (page - 1)).Limit(limit).All(&albums)
	if err != nil {
		return err, nil
	}
	return nil, albums
}

func (db *DB) GetComments(limit, page int, postID string) (error, []*m.Comment) {
	collection := db.Db.C(commentCollection)
	var comments []*m.Comment
	err := collection.Find(bson.M{"post_id": postID}).Skip(limit * (page - 1)).Limit(limit).All(&comments)
	if err != nil {
		return err, nil
	}
	return nil, comments
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

func (db *DB) GetPosts(pIDs []string) (error, []*m.Post) {
	collection := db.Db.C(postCollection)
	var posts []*m.Post
	listQueriesID := make([]bson.M, 0)
	for _, p := range pIDs {
		if p == "" {
			continue
		}
		bsonID := bson.M{"_id": bson.ObjectIdHex(p)}
		listQueriesID = append(listQueriesID, bsonID)
	}
	err := collection.Find(bson.M{"$or": listQueriesID}).All(&posts)
	if err != nil {
		return err, nil
	}
	return nil, posts
}

func (db *DB) GetUserOwns(uIDs []string) (error, []*m.User) {
	collection := db.Db.C(userCollection)
	var users []*m.User
	listQueriesID := make([]bson.M, 0)
	for _, p := range uIDs {
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
