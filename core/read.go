package core

import (
	_db "github.com/my0sot1s/social/db"
	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
)

/*
     +----------------------+       +------------------+
     | Find user following  |+--->+ | Create new Feed  |
+----+--------+-------------+       +------------------+
| Create Post |
+----+--------+

*/

type IRead interface {
	GetFeed(limit, page int, userID string) (error, []*m.Feed)
	GetPost(limit, page int, userID string) (error, []*m.Post)
	GetPostById(postID string) (error, *m.Post)
	GetComments(limit, page int, postID string) (error, []*m.Comment)
}

func (p *Core) Config(db *_db.DB) {
	p.Db = db
}

func (p *Core) GetPostByUser(limit, page int, user string) (error, []*m.Post) {
	err, posts := p.Db.GetPost(limit, page, user)
	for _, value := range posts {
		utils.Log(value.ID)
	}
	if err != nil {
		return err, nil
	}
	return nil, posts
}

func (p *Core) GetFeedByUser(limit, page int, user string) (error, []*m.Feed) {
	err, feeds := p.Db.GetFeed(limit, page, user)
	for _, value := range feeds {
		utils.Log(value.ID)
	}
	if err != nil {
		return err, nil
	}
	return nil, feeds
}

func (p *Core) GetPostID(pid string) (error, *m.Post) {
	err, posts := p.Db.GetPostById(pid)
	if err != nil {
		return err, nil
	}
	return nil, posts
}

func (p *Core) GetPostUser(limit, page int, userID string) (error, []*m.Post) {
	err, posts := p.Db.GetPost(limit, page, userID)
	if err != nil {
		return err, nil
	}
	return nil, posts
}

func (p *Core) GetCountLike(postId string) (error, int) {
	err, count := p.Db.CountLike(postId)
	if err != nil {
		return err, 0
	}
	return nil, count
}

func (p *Core) GetCommentByPostID(limit, page int, postId string) (error, []*m.Comment) {
	err, comments := p.Db.GetComments(limit, page, postId)
	if err != nil {
		return err, nil
	}
	return nil, comments
}

func (p *Core) GetFollowerByOwner(limit, page int, owner string) (error, []*m.User) {
	err, follower := p.Db.GetFollower(limit, page, owner)
	listUserID := make([]string, 0)
	for _, f := range follower {
		listUserID = append(listUserID, f.Own)
	}
	if err != nil {
		return err, nil
	}
	// get List user by 1 request
	users := make([]*m.User, 0)
	return nil, users
}
