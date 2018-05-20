package core

import (
	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
)

func (p *Social) LoadFeedByUser(limit int, anchor, userID string) (error, []*m.Feed, string) {
	err, feeds := p.Db.GetFeed(limit, anchor, userID)
	if err != nil {
		return err, nil, ""
	}
	var newAnchor string
	if len(feeds) > 0 {
		if limit > 0 {
			newAnchor = feeds[0].GetID()
		} else {
			newAnchor = feeds[len(feeds)-1].GetID()
		}
	}
	return nil, feeds, newAnchor
}

func (p *Social) LoadPostsByFeedUser(limit int, anchor, userID string) (error, []*m.Post, string) {
	err, feeds := p.Db.GetFeed(limit, anchor, userID)
	if err != nil {
		utils.ErrLog(err)
		return err, nil, ""
	}
	pIDs := make([]string, 0)

	for _, val := range feeds {
		pIDs = append(pIDs, val.GetPostID())
	}
	var err2 error
	posts := make([]*m.Post, 0)

	if err2, posts = p.Db.GetPosts(pIDs); err2 != nil {
		utils.ErrLog(err2)
	}
	newAnchor := ""
	if len(feeds) > 0 {
		if limit > 0 {
			newAnchor = feeds[0].GetID()
		} else {
			newAnchor = feeds[len(feeds)-1].GetID()
		}
	}
	return nil, posts, newAnchor
}
