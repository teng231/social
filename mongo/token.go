package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type AccessToken struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Token   string        `json:"token" bson:"token"`
	Created time.Time     `json:"created" bson:"created"`
	UserID  string        `json:"user_id" bson:"user_id"`
	Scopes  string        `json:"scopes" bson:"scopes"`
}

func (p *AccessToken) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *AccessToken) GetToken() string {
	if p.Token == "" {
		return ""
	}
	return p.Token
}
func (p *AccessToken) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}

func (p *AccessToken) GetUserID() string {
	if p.UserID == "" {
		return ""
	}
	return p.UserID
}
func (p *AccessToken) GetScopes() string {
	if p.Scopes == "" {
		return ""
	}
	return p.Scopes
}
