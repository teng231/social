package db

import "gopkg.in/mgo.v2/bson"

func (db *DB) CreatePost(p *Post) *Post {
	collection := db.db.C(postCollection)
	p.ID = bson.NewObjectId()
	err := collection.Insert(&p)
	if err != nil {
		panic(err)
	}
	return p
}

func (db *DB) CreateComment(c *Comment) *Comment {
	collection := db.db.C(commentCollection)
	c.ID = bson.NewObjectId()
	err := collection.Insert(&c)
	if err != nil {
		panic(err)
	}
	return c
}

func (db *DB) CreateFeed(f *Feed) *Feed {
	collection := db.db.C(feedCollection)
	f.ID = bson.NewObjectId()
	err := collection.Insert(&f)
	if err != nil {
		panic(err)
	}
	return f
}

func (db *DB) CreateUser(u *User) *User {
	collection := db.db.C(userCollection)
	u.ID = bson.NewObjectId()
	err := collection.Insert(&u)
	if err != nil {
		panic(err)
	}
	return u
}

func (db *DB) CreateToken(t *AccessToken) *AccessToken {
	collection := db.db.C(tokenCollection)
	t.ID = bson.NewObjectId()
	err := collection.Insert(&t)
	if err != nil {
		panic(err)
	}
	return t
}

func (db *DB) ModifyLike(t *Like, postID, userID string) *Like {
	collection := db.db.C(tokenCollection)
	t.ID = bson.NewObjectId()
	updator := bson.M{"$set", t}
	err := collection.Upsert(bson.M{"post_id": postID, "user_id": userID}, updator)
	if err != nil {
		panic(err)
	}
	return t
}

func (db *DB) ModifyFollower(t *Follower) *Follower {
	collection := db.db.C(tokenCollection)
	t.ID = bson.NewObjectId()
	err := collection.Insert(&t)
	if err != nil {
		panic(err)
	}
	return t
}

func (db *DB) CreateAlbum(a *Album) *Album {
	collection := db.db.C(albumCollection)
	a.ID = bson.NewObjectId()
	err := collection.Insert(&a)
	if err != nil {
		panic(err)
	}
	return a
}
