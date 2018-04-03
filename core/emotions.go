package core

import (
	"encoding/json"

	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
)

func (c *Social) CreateEmotion(media []*m.Media, by string, created int) []byte {
	emo := &m.Emotion{
		Medias:  media,
		By:      by,
		Created: created,
	}
	data, err := json.Marshal(emo)
	utils.ErrLog(err)
	return data
}
