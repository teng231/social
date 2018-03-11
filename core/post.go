package core

import (
	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
)

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

func (p *Core) GetPostID(pid string) (error, *m.Post) {
	err, post := p.Db.GetPostById(pid)
	if err != nil {
		return err, nil
	}
	return nil, post
}

func (p *Core) GetPostUser(limit, page int, userID string) (error, []*m.Post) {
	err, posts := p.Db.GetPost(limit, page, userID)
	if err != nil {
		return err, nil
	}
	return nil, posts
}
