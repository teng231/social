package core

import (
	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
)

func (p *Core) GetCountLike(postId string) (error, int) {
	err, count := p.Db.CountLike(postId)
	if err != nil {
		utils.ErrLog(err)
		return err, 0
	}
	return nil, count
}

func (p *Core) HitLikePost(postId, uID string) error {
	err := p.Db.HitLikePost(postId, uID)
	if err != nil {
		utils.ErrLog(err)
		return err
	}
	return nil
}

func (p *Core) UnlikePost(postId, uID string) error {
	err := p.Db.UnlikePost(postId, uID)
	if err != nil {
		utils.ErrLog(err)
		return err
	}
	return nil
}

func (p *Core) GetUserLikePost(postId string) (error, []*m.User) {
	// get like
	err1, liked := p.Db.GetLikes(postId)
	if err1 != nil {
		utils.ErrLog(err1)
		return err1, nil
	}
	uIDs := make([]string, 0)
	for _, v := range liked {
		uIDs = append(uIDs, v.GetUserID())
	}
	err, users := p.Db.GetUsersLikePost(uIDs)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	for _, v := range users {
		v.SetPassword("")
	}
	return nil, users
}
