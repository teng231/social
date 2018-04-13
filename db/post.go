package db

import (
	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetPost(limit int, anchor, userId string) (error, []*m.Post) {
	posts := make([]*m.Post, 0)
	err, ma := db.ReadByIdCondition(postCollection, anchor, limit, bson.M{"user_id": userId})
	if err != nil {
		return err, nil
	}
	for _, v := range ma {
		a := &m.Post{}
		a.ToPost(v)
		posts = append(posts, a)
	}
	return nil, posts
}

func (db *DB) GetExplore(limit int, anchor string, listIgnore []string) (error, []*m.Post) {
	posts := make([]*m.Post, 0)
	conditions := bson.M{"user_id": bson.M{"$nin": listIgnore}}
	err, ma := db.ReadByIdCondition(postCollection, anchor, limit, conditions)
	if err != nil {
		return err, nil
	}
	for _, v := range ma {
		a := &m.Post{}
		a.ToPost(v)
		posts = append(posts, a)
	}
	return nil, posts
}

func (db *DB) GetPostById(postID string) (error, *m.Post) {
	var post = &m.Post{}
	err, i := db.ReadById(postCollection, postID)
	if err != nil {
		return err, nil
	}
	utils.Log(i)
	post.ToPost(i)
	return nil, post
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
	utils.Log(pIDs)
	err := collection.Find(bson.M{"$or": listQueriesID}).Sort("-created").All(&posts)
	for _, p := range posts {
		utils.Log(p)
	}

	if err != nil {
		return err, nil
	}
	return nil, posts
}

func (db *DB) CreatePost(p *m.Post) error {
	collection := db.Db.C(postCollection)
	p.ID = bson.NewObjectId()
	err := collection.Insert(&p)
	if err != nil {
		return err
	}
	return nil
}
