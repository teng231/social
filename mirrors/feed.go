package mirror

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

type Feed struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ConsumerID string        `json:"comsumer_id" bson:"comsumer_id"`
	PostID     string        `json:"post_id" bson:"post_id"`
	Author     string        `json:"author" bson:"author"`
	Created    time.Time     `json:"created" bson:"created"`
}

func (p *Feed) ToFeed(m map[string]interface{}) {
	m["id"] = m["_id"]
	str, err := json.Marshal(m)
	utils.ErrLog(err)
	json.Unmarshal(str, &p)
}

func (p *Feed) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
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
func (p *Feed) GetAuthor() string {
	if p.Author == "" {
		return ""
	}
	return p.Author
}

func (p *Feed) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}
