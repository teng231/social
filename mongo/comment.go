package mongo

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PostID  string        `json:"post_id" bson:"post_id"`
	UserID  string        `json:"user_id" bson:"user_id"`
	Text    string        `json:"text" bson:"text"`
	Created time.Time     `json:"created" bson:"created"`
}

func (p *Comment) ToComment(m map[string]interface{}) {
	m["id"] = m["_id"]
	str, err := json.Marshal(m)
	utils.ErrLog(err)
	json.Unmarshal(str, &p)
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
