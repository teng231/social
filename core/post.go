package core

import (
	"errors"
	"time"

	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
)

func (p *Social) LoadPostID(pid string) (error, *m.Post) {
	err, post := p.Db.GetPostById(pid)
	if err != nil {
		return err, nil
	}
	return nil, post
}

func (p *Social) LoadPostUser(limit int, anchor, userID string) (error, []*m.Post, string) {
	err, posts := p.Db.GetPost(limit, anchor, userID)
	if err != nil {
		return err, nil, ""
	}
	newAnchor := ""
	if len(posts) > 0 {
		if limit > 0 {
			newAnchor = posts[0].GetID()
		} else {
			newAnchor = posts[len(posts)-1].GetID()
		}
	}
	return nil, posts, newAnchor
}

// AddNewPost `uid` owner of post
func (c *Social) AddNewPostBonusFeed(userID, content, mediasStr, tagsStr string) (error, *m.Post) {
	// create post
	var medias []*m.Media
	errMedia := utils.Str2T(mediasStr, &medias)
	if errMedia != nil {
		return errors.New("Media not valid"), nil
	}
	var tags []string
	if tagsStr != "" {
		errTags := utils.Str2T(tagsStr, &tags)
		if errTags != nil {
			return errors.New("Tags not valid"), nil
		}
	} else {
		tags = make([]string, 0)
	}
	post := &m.Post{
		Created: time.Now(),
		UserID:  userID,
		Text:    content,
		Media:   medias,
		Tags:    tags,
	}
	owner := post.GetUserID()
	if owner == "" {
		return errors.New("no owner"), nil
	}
	if len(post.GetMedia()) == 0 {
		return errors.New("no media found"), nil
	}
	post.Created = time.Now()
	// create post
	err := c.Db.CreatePost(post)
	if err != nil {
		return err, nil
	}
	// find all user follow own
	err, follower := c.Db.GetFollower(owner)
	feeds := make([]*m.Feed, 1)
	feeds[0] = &m.Feed{
		Created:    time.Now(),
		ConsumerID: owner,
		PostID:     post.GetID(),
	}
	for _, v := range follower {
		peopleFollowOwner := v.GetFollower()
		if peopleFollowOwner == "" {
			continue
		}
		feed := &m.Feed{
			Created:    time.Now(),
			ConsumerID: v.GetFollower(),
			PostID:     post.GetID(),
		}
		feeds = append(feeds, feed)
	}
	err2, _ := c.Db.CreateFeeds(feeds)
	if err2 != nil {
		return err2, nil
	}
	return nil, post
}

func (p *Social) GetAnyPost(limit int, anchor, owner string) (error, []*m.Post, string) {
	// get all follower
	var err error
	err, listFollowing := p.Db.GetFollowing(owner)
	listUid := make([]string, 0)
	listUid = append(listUid, owner)
	for _, v := range listFollowing {
		listUid = append(listUid, v.GetOwn())
	}
	err, posts := p.Db.GetExplore(limit, anchor, listUid)
	if err != nil {
		return err, nil, ""
	}
	newAnchor := ""
	if len(posts) > 0 {
		if limit > 0 {
			newAnchor = posts[0].GetID()
		} else {
			newAnchor = posts[len(posts)-1].GetID()
		}
	}
	return nil, posts, newAnchor
}
