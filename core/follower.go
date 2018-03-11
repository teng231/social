package core

import m "github.com/my0sot1s/social/mongo"

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
