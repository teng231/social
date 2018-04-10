package core

import (
	"time"

	m "github.com/my0sot1s/social/mirrors"
)

func (c *Social) CreateUserSave(uid, pid string) (error, *m.Saved) {
	saved := &m.Saved{
		PostId:  pid,
		Saver:   uid,
		Created: time.Now(),
	}
	err := c.Db.CreateSaved(saved)
	if err != nil {
		return err, nil
	}
	return nil, saved
}

func (c *Social) ListUserSaved(limit int, anchor, uid string) (error, string, []*m.Saved) {
	err, listSaved := c.Db.ListSaved(limit, anchor, uid)
	if err != nil {
		return err, "", nil
	}
	newanchor := ""
	if len(listSaved) > 0 {
		newanchor = listSaved[len(listSaved)-1].GetID()
	}

	return nil, newanchor, listSaved
}
