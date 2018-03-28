package core

import (
	"time"

	m "github.com/my0sot1s/social/mongo"
)

func (c *Core) CreateUserSave(uid, pid string) (error, *m.Saved) {
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
