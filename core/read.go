package core

import (
	_db "social/db"
	m "social/mongo"
	"social/utils"
)

/*
     +----------------------+       +------------------+
     | Find user following  |+--->+ | Create new Feed  |
+----+--------+-------------+       +------------------+
| Create Post |
+----+--------+

*/

type IRead interface {
	GetFeed(limit, page int, userID string) []*m.Feed
	GetPost(limit, page int, userID string) []*m.Post
	GetPostById(postID string) *m.Post
}

func (p *Core) Config(db *_db.DB) {
	p.Db = db
}

func (p *Core) GetPostByUser(limit, page int, user string) []*m.Post {
	posts := p.Db.GetPost(limit, page, user)
	for _, value := range posts {
		utils.Log(value.ID)
	}
	return posts
}

func (p *Core) GetFeedByUser(limit, page int, user string) []*m.Feed {
	feeds := p.Db.GetFeed(limit, page, user)
	for _, value := range feeds {
		utils.Log(value.ID)
	}
	return feeds
}

func (p *Core) GetPostID(user string) *m.Post {
	posts := p.Db.GetPostById(user)
	return posts
}

func (p *Core) GetCountLike(postId string) int {
	count := p.Db.CountLike(postId)
	return count
}

func (p *Core) GetFollowerByOwner(limit, page int, owner string) []*m.User {
	follower := p.Db.GetFollower(limit, page, owner)
	listUserId := make([]string, 0)
	for _, f := range follower {
		listUserId = append(listUserId, f.Own)
	}
	// get List user by 1 request
	users := make([]*m.User, 0)
	return users
}
