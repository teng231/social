package mirror

import (
	"time"
)

type Emotion struct {
	Medias  []*Media
	Created int
	By      string
}

func (e *Emotion) GetMedia() []*Media {
	if &e.Medias == nil {
		return make([]*Media, 0)
	}
	return e.Medias
}

func (e *Emotion) GetCreated() int {
	if &e.Created == nil {
		t := time.Now()
		return int(t.Unix())
	}
	return e.Created
}

func (e *Emotion) GetBy() string {
	if &e.By == nil {
		return ""
	}
	return e.By
}
