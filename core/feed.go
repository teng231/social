package core

import (
	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
)

func (p *Core) LoadFeedByUser(limit, page int, userID string) (error, []*m.Feed) {
	err, feeds := p.Db.GetFeed(limit, page, userID)
	if err != nil {
		return err, nil
	}
	return nil, feeds
}

func (p *Core) LoadPostsByFeedUser(limit, page int, userID string) (error, []*m.Post, []*m.User) {
	err, feeds := p.Db.GetFeed(limit, page, userID)
	if err != nil {
		utils.ErrLog(err)
		return err, nil, nil
	}
	pIDs := make([]string, 0)

	for _, val := range feeds {
		pIDs = append(pIDs, val.GetPostID())
	}
	uIDs := make([]string, 0)
	var err2 error
	var posts []*m.Post
	var users []*m.User
	err2, posts = p.Db.GetPosts(pIDs)
	if err2 != nil {
		utils.ErrLog(err2)
	}
	for _, v := range posts {
		uIDs = append(uIDs, v.GetUserID())
	}
	err2, users = p.getUserByIDs(uIDs)
	if err2 != nil {
		utils.ErrLog(err2)
	}
	return nil, posts, users
}
