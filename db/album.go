package db

import (
	m "github.com/my0sot1s/social/mirrors"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetAlbum(AlbumID string) (error, *m.Album) {
	var album = &m.Album{}
	err, i := db.ReadById(albumCollection, AlbumID)
	if err != nil {
		return err, nil
	}
	album.ToAlbum(i)
	return nil, album
}

func (db *DB) GetAlbumByAuthor(limit int, anchor, userId string) (error, []*m.Album) {
	albums := make([]*m.Album, 0)
	err, ma := db.ReadByIdCondition("_id", albumCollection, anchor, limit, bson.M{"author": userId})
	if err != nil {
		return err, nil
	}
	for _, v := range ma {
		a := &m.Album{}
		a.ToAlbum(v)
		albums = append(albums, a)
	}
	return nil, albums
}

func (db *DB) CreateAlbum(a *m.Album) error {
	collection := db.Db.C(albumCollection)
	a.ID = bson.NewObjectId()
	err := collection.Insert(&a)
	if err != nil {
		return err
	}
	return nil
}
