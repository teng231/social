package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PostID  string        `json:"post_id" bson:"post_id"`
	UserID  string        `json:"user_id" bson:"user_id"`
	Text    string        `json:"text" bson:"content"`
	Created time.Time     `json:"created" bson:"created"`
}

func (p *Comment) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *Comment) GetPostID() string {
	if p.PostID == "" {
		return ""
	}
	return p.PostID
}
func (p *Comment) GetUserID() string {
	if p.UserID == "" {
		return ""
	}
	return p.UserID
}
func (p *Comment) GetText() string {
	if p.Text == "" {
		return ""
	}
	return p.Text
}
func (p *Comment) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}
