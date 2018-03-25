package core

import (
	"errors"
	"time"

	m "github.com/my0sot1s/social/mongo"
	"github.com/my0sot1s/social/utils"
)

func (p *Core) LoadCommentByPostID(limit int, anchor, postId string) (error, []*m.Comment) {
	err, comments := p.Db.GetComments(limit, anchor, postId)
	if err != nil {
		return err, nil
	}
	return nil, comments
}

// InsertPost add post to db
func (c *Core) UpsertCommentsToPost(pid, text, userID string) (error, *m.Comment) {
	if pid == "" || userID == "" {
		utils.ErrLog(errors.New("err no field pid or userID"))
		return errors.New("err no field pid or userID"), nil
	}
	cmt := &m.Comment{
		PostID:  pid,
		UserID:  userID,
		Created: time.Now(),
		Text:    text,
	}
	err, upCmt := c.Db.CreateComment(cmt)

	if err != nil {
		utils.ErrLog(err)
	}

	return err, upCmt
}
