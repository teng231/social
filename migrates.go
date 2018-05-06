package main

import (
	"sync"

	"github.com/my0sot1s/social/db"
	m "github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
)

func RunMigrateFeedToFollow() {
	// lấy dữ liệu của trang feeds
	// cứ postID get user_id
	// insert userId vào feed

	c := loadConfig()
	mg := &db.DB{}
	mg.Config(c.DbHost, c.DbName, c.Username, c.Password)
	_, feeds := mg.GetMigrateFeed(100)
	wg := sync.WaitGroup{}
	wg.Add(len(feeds))
	for _, v := range feeds {
		go func(pid string, f *m.Feed) {
			err, p := mg.GetPostById(pid)
			if err != nil {
				utils.ErrLog(err)
				utils.Log(pid)
				wg.Done()
				return
			}
			f.Author = p.GetUserID()
			mg.UpsertFeed(f.GetID(), f)
			wg.Done()
		}(v.GetPostID(), v)
	}
	wg.Wait()
}
