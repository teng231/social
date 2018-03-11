package core

import (
	dbase "github.com/my0sot1s/social/db"
	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/redis"
)

type Core struct {
	Db    *dbase.DB
	token *JWTAuthentication
	rd    *redis.RedisCli
}
type ICore interface {
	GetFeed(limit, page int, userID string) (error, []*m.Feed)
	GetPost(limit, page int, userID string) (error, []*m.Post)
	GetPostById(postID string) (error, *m.Post)
	GetComments(limit, page int, postID string) (error, []*m.Comment)
	CreatePost(*m.Post) *m.Post
	CreateFeeds(*m.Post) []*m.Post
}

func (c *Core) Config(db *dbase.DB, rd *redis.RedisCli, privateKeyPath, PublicKeyPath string) {
	// connect to drive Mongo
	c.Db = db
	// connect token access
	c.token = &JWTAuthentication{}
	c.token.Config(privateKeyPath, PublicKeyPath)
	// connect drive Redis
	c.rd = rd
}
