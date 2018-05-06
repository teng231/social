package core

import (
	dbase "github.com/my0sot1s/social/db"
	"github.com/my0sot1s/social/mail"
	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/redis"
)

type Social struct {
	HOST   string
	Db     ISocial
	token  *JWTAuthentication
	rd     *redis.RedisCli
	mailAd *mail.EmailMgr
}

type ISocial interface {
	GetAlbumByAuthor(limit int, anchor, userId string) (error, []*m.Album)

	GetUserByUname(username string) (error, *m.User)
	GetUserByEmail(email string) (error, *m.User)
	GetPost(limit int, anchor, userId string) (error, []*m.Post)
	GetPostById(postID string) (error, *m.Post)
	GetFeed(limit int, anchor, userId string) (error, []*m.Feed)
	GetMigrateFeed(limit int) (error, []*m.Feed)
	GetFollower(own string) (error, []*m.Follower)
	GetFollowing(follower string) (error, []*m.Follower)
	CountLike(postID string) (error, int)
	GetAlbum(AlbumID string) (error, *m.Album)
	GetComments(limit int, anchor, postID string) (error, []*m.Comment)
	GetLikes(postID string) (error, []*m.Like)
	IsUserLikePost(pid, uid string) (error, bool)
	GetPosts(pIDs []string) (error, []*m.Post)
	GetExplore(limit int, anchor string, listIgnore []string) (error, []*m.Post)
	GetUserByIds(uIDs []string) (error, []*m.User)
	GetUserById(id string) (error, *m.User)
	CreatePost(p *m.Post) error
	CreateComment(c *m.Comment) (error, *m.Comment)
	CountCommentByPostId(postID string) (error, int)
	CreateFeed(f *m.Feed) error
	CreateFeeds(feeds []*m.Feed) (error, []interface{})
	CreateUser(u *m.User) error
	UpsertFeed(id string, f *m.Feed) error

	ModifyFollower(t *m.Follower) (error, *m.Follower)
	CreateAlbum(a *m.Album) error
	HitLikePost(postID, userID string) error
	UnlikePost(postID, userID string) error
	// GetUsersLikePost(userIDs []string) (error, []*m.User)
	FollowUser(f *m.Follower) error
	UnfollowUser(own, uid string) error
	//
	UpdateStateUser(uid, state string) error
	// ReadById(cName, anyId string) (error, map[string]interface{})
	// ReadByIdCondition(cName, anchor string, limit int, conditions map[string]interface{}) (error, []map[string]interface{})
	UpdateUserPassword(uid, password string) error

	RemoveSaved(sid string) error
	ListSaved(limit int, anchor, uid string) (error, []*m.Saved)
	CreateSaved(saved *m.Saved) error
	CountFollower(own string) (error, int)
	IsFollow(own, uid string) (int, error)
}

func (c *Social) Config(host string, db *dbase.DB, rd *redis.RedisCli, mailAd *mail.EmailMgr, privateKeyPath, PublicKeyPath string) {
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

// func (c *Social) CoreTest() {}
