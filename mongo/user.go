package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Password  string        `json:"password" bson:"password"`
	AlbumName string        `json:"album_name" bson:"album_name`
	UserName  string        `json:"username" bson:"username"`
	Email     string        `json:"email" bson:"email"`
	Created   time.Time     `json:"created" bson:"created"`
	Avatar    string        `json:"avatar" bson:"avatar"`
	Banner    string        `json:"banner" bson:"banner"`
	State     string        `json:"state" bson:"state"`
}

func (p *User) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.String()
}

func (p *User) GetAlbumName() string {
	if p.AlbumName == "" {
		return ""
	}
	return p.AlbumName
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
