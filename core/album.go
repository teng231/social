package core

import (
	m "github.com/my0sot1s/social/mongo"
)

func (p *Core) LoadAlbumByAuthor(limit, page int, userID string) (error, []*m.Album) {
	err, albums := p.Db.GetAlbumByAuthor(limit, page, userID)
	if err != nil {
		return err, nil
	}
	return nil, albums
}

func (p *Core) LoadAlbumById(ID string) (error, *m.Album) {
	err, album := p.Db.GetAlbum(ID)
	if err != nil {
		return err, nil
	}
	return nil, album
}

func (p *Core) UpsertAnAlbum(a *m.Album) (error, *m.Album) {
	err, album := p.Db.CreateAlbum(a)
	if err != nil {
		return err, nil
	}
	return nil, album
}

// func (p *Core) ModifiMedia(albumID string, medias []*m.Media) (error, *m.Album) {
// 	err, album := p.Db.CreateAlbum(a)
// 	if err != nil {
// 		return err, nil
// 	}
// 	return nil, album
// }
