package core

import (
	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
)

func (p *Social) LoadCountLike(postId string) (error, int) {
	err, count := p.Db.CountLike(postId)
	if err != nil {
		utils.ErrLog(err)
		return err, 0
	}
	return nil, count
}

func (p *Social) UpsertLikePost(postId, uID string) error {
	err := p.Db.HitLikePost(postId, uID)
	if err != nil {
		utils.ErrLog(err)
		return err
	}
	return nil
}

func (p *Social) RemoveLikePost(postId, uID string) error {
	err := p.Db.UnlikePost(postId, uID)
	if err != nil {
		utils.ErrLog(err)
		return err
	}
	return nil
}

func (p *Social) LoadUserLikePost(postId string) (error, []*m.User) {
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
	err, users := p.Db.GetUserByIds(uIDs)
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	for _, v := range users {
		v.SetPassword("")
	}
	return nil, users
}

func (p *Social) CheckOwnerLikePost(postId, uid string) (error, bool) {
	// get like
	err, b := p.Db.IsUserLikePost(postId, uid)
	if err != nil && b {
		utils.ErrLog(err)
		return err, true
	}
	return err, b
}
