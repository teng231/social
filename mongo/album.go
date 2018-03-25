package mongo

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/social/utils"
	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AlbumName  string        `json:"album_name" bson:"album_name"`
	AuthorID   string        `json:"author" bson:"author"`
	AlbumMedia []Media       `json:"album_media" bson:"album_media"`
	Created    time.Time     `json:"created" bson:"created"`
	Modified   time.Time     `json:"modified,omitempty" bson:"modified,omitempty"`
}

func (p *Album) ToAlbum(m map[string]interface{}) {
	utils.Log(m)
	m["id"] = m["_id"]
	utils.Log(m)
	str, err := json.Marshal(m)
	utils.ErrLog(err)
	json.Unmarshal(str, &p)
}

func (p *Album) GetID() string {
	if !p.ID.Valid() {
		return ""
	}
	return p.ID.Hex()
}

func (p *Album) GetAlbumName() string {
	if p.AlbumName == "" {
		return ""
	}
	return p.AlbumName
}

func (p *Album) GetAuthorID() string {
	if p.AuthorID == "" {
		return ""
	}
	return p.AuthorID
}

func (p *Album) GetCreated() time.Time {
	if p.Created.IsZero() {
		return time.Now()
	}
	return p.Created
}

func (p *Album) GetModified() time.Time {
	if p.Modified.IsZero() {
		return time.Now()
	}
	return p.Modified
}
