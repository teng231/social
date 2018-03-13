package mongo

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Follower struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Own      string        `json:"own" bson:"own"`
	Follower string        `json:"follower" bson:"follower"`
	Created  time.Time     `json:"created" bson:"created"`
	State    bool          `json:"state" bson:"state"`
}

func (p *Follower) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *Follower) GetOwn() string {
	if p.Own == "" {
		return ""
	}
	return p.Own
}

func (p *Follower) GetFollower() string {
	if p.Follower == "" {
		return ""
	}
	return p.Follower
}

func (p *Follower) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}

func (p *Follower) GetState() bool {
	if p.State != true && p.State != false {
		return false
	}
	return true
}
