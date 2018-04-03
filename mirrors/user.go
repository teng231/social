package mirror

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Password  string        `json:"password,omitempty" bson:"password"`
	AlbumName string        `json:"albumname" bson:"albumname"`
	UserName  string        `json:"username" bson:"username"`
	Email     string        `json:"email" bson:"email"`
	Created   time.Time     `json:"created" bson:"created"`
	Avatar    string        `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Banner    string        `json:"banner,omitempty" bson:"banner,omitempty"`
	State     string        `json:"state" bson:"state"`
}

func (p *User) ToUser(m map[string]interface{}) {
	m["id"] = m["_id"]
	str, err := json.Marshal(m)
	utils.ErrLog(err)
	json.Unmarshal(str, &p)
}
func (p *User) SetPassword(newPw string) {
	p.Password = newPw
}

func (p *User) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *User) GetAlbumName() string {
	if p.AlbumName == "" {
		return ""
	}
	return p.AlbumName
}

func (p *User) GetPassword() string {
	if p.Password == "" {
		return ""
	}
	return p.Password
}

func (p *User) GetUserName() string {
	if p.UserName == "" {
		return ""
	}
	return p.UserName
}

func (p *User) GetEmail() string {
	if p.Email == "" {
		return ""
	}
	return p.Email
}

func (p *User) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}

func (p *User) GetAvatar() string {
	if p.Avatar == "" {
		return ""
	}
	return p.Avatar
}

func (p *User) GetBanner() string {
	if p.Banner == "" {
		return ""
	}
	return p.Banner
}

func (p *User) GetState() string {
	if p.State == "" {
		return "inactived"
	}
	return p.State
}
