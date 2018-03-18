package core

import (
	"fmt"
	"time"

	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
)

func (p *Core) LoadAlbumByAuthor(limit, page int, userID string) (error, []*m.Album) {
	err, albums := p.Db.GetAlbumByAuthor(limit, page, userID)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	return nil, albums
}

func (p *Core) LoadAlbumById(ID string) (error, *m.Album) {
	err, album := p.Db.GetAlbum(ID)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	return nil, album
}

func (p *Core) UpsertAnAlbum(albumName, media, owner string) (error, *m.Album) {
	if albumName == "" {
		albumName = fmt.Sprintf("created-%d", time.Now().Second())
	}
	var mediaArray []m.Media
	err := utils.Str2T(media, &mediaArray)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	newAlbum := &m.Album{
		AuthorID:   owner,
		AlbumName:  albumName,
		Created:    time.Now(),
		AlbumMedia: mediaArray,
	}
	err, album := p.Db.CreateAlbum(newAlbum)
	if err != nil {
		utils.ErrLog(err)
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
