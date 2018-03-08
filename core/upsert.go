package core

import (
	m "social/mongo"
)

type CInsert interface {
	CreatePost(*m.Post) *m.Post
	CreateFeeds(*m.Post) []*m.Post
}

// InsertPost add post to db
func (c *Core) InsertPost(raw map[string]interface{}) {
	// postData := &m.Post{}
	// post := c.Db.CreatePost(postData)
}
