package mongo

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

// Post define
type Post struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID   string        `json:"user_id" bson:"user_id"`
	Text     string        `json:"text" bson:"text"`
	Created  time.Time     `json:"created" bson:"created"`
	Modified time.Time     `json:"modified,omitempty" bson:"modified,omitempty"`
	Media    []*Media      `json:"media" bson:"media"`
	Tags     []string      `json:"tags,omitempty" bson:"tags"`
}

func (p *Post) ToPost(m map[string]interface{}) {
	m["id"] = m["_id"]
	str, err := json.Marshal(m)
	utils.ErrLog(err)
	json.Unmarshal(str, &p)
}
func (p *Post) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *Post) GetUserID() string {
	if p.UserID == "" {
		return ""
	}
	return p.UserID
}

func (p *Post) GetText() string {
	if p.Text == "" {
		return ""
	}
	return p.Text
}

func (p *Post) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}

func (p *Post) GetModified() time.Time {
	if p.Modified.IsZero() {
		return time.Now()
	}
	return p.Modified
}

func (p *Post) GetMedia() []*Media {
	if len(p.Media) == 0 {
		return make([]*Media, 0)
	}
	return p.Media
}

func (p *Post) GetTags() []string {
	if len(p.Tags) == 0 {
		return make([]string, 0)
	}
	return p.Tags
}
