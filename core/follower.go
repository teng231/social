package core

import (
	"time"

	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
)

//GetFollowerByOwner  ai đang follow bạn `owner`
func (p *Core) GetFollowerByOwner(owner string) (error, []*m.User) {
	err, follower := p.Db.GetFollower(owner)
	listUserID := make([]string, 0)
	for _, f := range follower {
		listUserID = append(listUserID, f.GetFollower())
	}
	if len(listUserID) == 0 {
		return nil, nil
	}
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}
	err2, users := p.Db.GetUserOwns(listUserID)
	// get List user by 1 request
	if err2 != nil {
		utils.ErrLog(err2)
		return err2, nil
	}
	return nil, users
}

// GetFollowingByUid bạn đang follow những ai `userId`
func (p *Core) GetFollowingByUid(uid string) (error, []*m.User) {
	err, follower := p.Db.GetFollowing(uid)
	listUserID := make([]string, 0)
	for _, f := range follower {
		listUserID = append(listUserID, f.GetOwn())
	}
	if len(listUserID) == 0 {
		return nil, nil
	}
	if err != nil {
		utils.ErrLog(err)
		return err, nil
	}

	err2, users := p.Db.GetUserOwns(listUserID)
	if err2 != nil {
		utils.ErrLog(err2)
		return err2, nil
	}
	return nil, users
}

func (p *Core) FollowAnUser(uid, owner string) error {
	follow := &m.Follower{
		Follower: uid,
		Own:      owner,
		Created:  time.Now(),
	}
	err := p.Db.FollowUser(follow)
	if err != nil {
		return err
	}
	return nil
}

func (p *Core) UnfollowAnUser(uid, owner string) error {
	err := p.Db.UnfollowUser(owner, uid)
	if err != nil {
		return err
	}
	return nil
}
