package mongo

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

type Saved struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PostId  string        `json:"post_id" bson:"post_id"`
	Saver   string        `json:"saver" bson:"saver"`
	Created time.Time     `json:"created" bson:"created"`
}

func (p *Saved) ToSaved(m map[string]interface{}) {
	m["id"] = m["_id"]
	str, err := json.Marshal(m)
	utils.ErrLog(err)
	json.Unmarshal(str, &p)
}

func (p *Saved) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *Saved) GetPostId() string {
	if p.PostId == "" {
		return ""
	}
	return p.PostId
}

func (p *Saved) GetSaver() string {
	if p.Saver == "" {
		return ""
	}
	return p.Saver
}
