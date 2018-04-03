package core

import (
	"fmt"
	"time"

	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
)

func (p *Social) LoadAlbumByAuthor(limit int, anchor, userID string) (error, []*m.Album) {
	err, albums := p.Db.GetAlbumByAuthor(limit, anchor, userID)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	return nil, albums
}

func (p *Social) LoadAlbumById(ID string) (error, *m.Album) {
	err, album := p.Db.GetAlbum(ID)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	return nil, album
}

func (p *Social) UpsertAnAlbum(albumName, media, owner string) (error, *m.Album) {
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
	err = p.Db.CreateAlbum(newAlbum)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	return nil, newAlbum
}

// func (p *Social) ModifiMedia(albumID string, medias []*m.Media) (error, *m.Album) {
// 	err, album := p.Db.CreateAlbum(a)
// 	if err != nil {
// 		return err, nil
// 	}
// 	return nil, album
// }
