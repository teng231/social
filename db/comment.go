package db

import (
	m "github.com/my0sot1s/social/mirrors"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetComments(limit int, anchor, postID string) (error, []*m.Comment) {
	comments := make([]*m.Comment, 0)
	err, mc := db.ReadByIdCondition("_id", commentCollection, anchor, limit, bson.M{"post_id": postID})
	if err != nil {
		return err, nil
	}
	for _, v := range mc {
		c := &m.Comment{}
		c.ToComment(v)
		comments = append(comments, c)
	}
	return nil, comments
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

func (db *DB) CountCommentByPostId(postID string) (error, int) {
	collection := db.Db.C(commentCollection)
	count, err := collection.Find(bson.M{"post_id": postID}).Count()
	if err != nil {
		return err, 0
	}
	return nil, count
}
