package mongo

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

type Like struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PostID  string        `json:"post_id" bson:"post_id"`
	UserID  string        `json:"user_id" bson:"user_id"`
	Created time.Time     `json:"created" bson:"created"`
	State   bool          `json:"state" bson:"state"`
}

func (p *Like) ToLike(m map[string]interface{}) {
	m["id"] = m["_id"]
	str, err := json.Marshal(m)
	utils.ErrLog(err)
	json.Unmarshal(str, &p)
}

func (p *Like) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *Like) GetPostID() string {
	if p.PostID == "" {
		return ""
	}
	return p.PostID
}

func (p *Like) GetUserID() string {
	if p.UserID == "" {
		return ""
	}
	return p.UserID
}

func (p *Like) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}

func (p *Like) GetState() bool {
	if p.State != true && p.State != false {
		return false
	}
	return true
}
