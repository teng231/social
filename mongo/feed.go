package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Feed struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ConsumerID string        `json:"consumer_id bson:"consumer_id"`
	PostID     string        `json:"post_id" bson:"post_id"`
	Created    time.Time     `json:"created" bson:"created"`
}

func (p *Feed) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.String()
}

func (p *Feed) GetConsumerID() string {
	if p.ConsumerID == "" {
		return ""
	}
	return p.ConsumerID
}

func (p *Feed) GetPostID() string {
	if p.PostID == "" {
		return ""
	}
	return p.PostID
}

func (p *Feed) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}
