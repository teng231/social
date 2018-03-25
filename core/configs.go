package core

import (
	dbase "github.com/my0sot1s/social/db"
	"github.com/my0sot1s/social/mail"
	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/redis"
)

type Core struct {
	HOST   string
	Db     ICore
	token  *JWTAuthentication
	rd     *redis.RedisCli
	mailAd *mail.EmailMgr
}

type ICore interface {
	GetAlbumByAuthor(limit int, anchor, userId string) (error, []*m.Album)

	GetUserByUname(username string) (error, *m.User)
	GetUserByEmail(email string) (error, *m.User)
	GetPost(limit int, anchor, userId string) (error, []*m.Post)
	GetPostById(postID string) (error, *m.Post)
	GetFeed(limit int, anchor, userId string) (error, []*m.Feed)
	GetFollower(own string) (error, []*m.Follower)
	GetFollowing(follower string) (error, []*m.Follower)
	CountLike(postID string) (error, int)
	GetAlbum(AlbumID string) (error, *m.Album)
	GetComments(limit int, anchor, postID string) (error, []*m.Comment)
	GetLikes(postID string) (error, []*m.Like)
	IsUserLikePost(pid, uid string) (error, bool)
	GetPosts(pIDs []string) (error, []*m.Post)
	GetUserByIds(uIDs []string) (error, []*m.User)
	CreatePost(p *m.Post) (error, *m.Post)
	CreateComment(c *m.Comment) (error, *m.Comment)
	CreateFeed(f *m.Feed) (error, *m.Feed)
	CreateFeeds(feeds []*m.Feed) (error, []interface{})
	CreateUser(u *m.User) (error, *m.User)
	ModifyFollower(t *m.Follower) (error, *m.Follower)
	CreateAlbum(a *m.Album) (error, *m.Album)
	HitLikePost(postID, userID string) error
	UnlikePost(postID, userID string) error
	//
	// GetUsersLikePost(userIDs []string) (error, []*m.User)
	FollowUser(f *m.Follower) error
	UnfollowUser(own, uid string) error
	//
	UpdateStateUser(uid, state string) error
	// ReadById(cName, anyId string) (error, map[string]interface{})
	// ReadByIdCondition(cName, anchor string, limit int, conditions map[string]interface{}) (error, []map[string]interface{})
	UpdateUserPassword(uid, password string) error
}

func (c *Core) Config(host string, db *dbase.DB, rd *redis.RedisCli, mailAd *mail.EmailMgr, privateKeyPath, PublicKeyPath string) {
	c.HOST = host
	// connect to drive Mongo
	c.Db = db
	// connect token access
	c.token = &JWTAuthentication{}
	c.token.Config(privateKeyPath, PublicKeyPath)
	// connect drive Redis
	c.rd = rd
	// connect mail adapters
	c.mailAd = mailAd
}

func (c *Core) CoreTest() {}
