package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Media define
type Media struct {
	PublicID string `json:"public_id" bson:"public_id"`
	Width    int32  `json:"width" bson:"width"`
	Height   int32  `json:"height" bson:"height"`
	Format   string `json:"format" bson:"format"`
	Bytes    string `json:"bytes" bson:"bytes"`
	URL      string `json:"byte" bson:"url"`
}

// Post define
type Post struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID   string        `json:"user_id" bson:"user_id"`
	Text     string        `json:"text" bson:"post_content"`
	Created  time.Time     `json:"created" bson:"created"`
	Modified time.Time     `json:"modified" bson:"modified"`
	Media    []*Media      `json:"media" bson:"media"`
	Tags     []string      `json:"tags" bson:"tags"`
}

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Password  string        `json:"password" bson:"password"`
	AlbumName string        `json:"album_name" bson:"album_name`
	UserName  string        `json:"username" bson:"username"`
	Email     string        `json:"email" bson:"email"`
	Created   time.Time     `json:"created" bson:"created"`
	Avatar    string        `json:"avatar" bson:"avatar"`
	Banner    string        `json:"banner" bson:"banner"`
}

type Album struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AlbumName  string        `json:"album_name" bson:"album_name`
	AuthorID   string        `json:"author" bson:"author"`
	AlbumMedia []*Media      `json:"media" bson:"album_media"`
	Created    time.Time     `json:"created" bson:"created"`
	Modified   time.Time     `json:"modified" bson:"modified"`
}

type Comment struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PostID  bson.ObjectId `json:"post_id" bson:"post_id"`
	UserID  string        `json:"user_id" bson:"user_id"`
	Text    string        `json:"text" bson:"content"`
	Created time.Time     `json:"created" bson:"created"`
}

type Follower struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Own      string        `json:"own" bson:"own"`
	Follower string        `json:"follower bson:"follower"`
	Created  time.Time     `json:"created" bson:"created"`
	State    bool          `json:"state" bson:"state"`
}

type Like struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PostID  bson.ObjectId `json:"post_id" bson:"post_id"`
	UserID  string        `json:"user_id" bson:"user_id"`
	Created time.Time     `json:"created" bson:"created"`
	State   bool          `json:"state" bson:"state"`
}

type Feed struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ConsumerID string        `json:"consumer_id bson:"consumer_id"`
	PostID     string        `json:"post_id" bson:"post_id"`
	Created    time.Time     `json:"created" bson:"created"`
}

type AccessToken struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Token   string        `json:"token" bson:"token"`
	Created time.Time     `json:"created" bson:"created"`
	UserID  string        `json:"user_id" bson:"user_id"`
	Scopes  string        `json:"scopes" bson:"scopes"`
}
