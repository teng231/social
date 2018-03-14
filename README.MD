# Core system: Backend golang for social networking

<center>
  <img src='./social.jpg'/>
  </br>
</center>

### Quy tắc
* Chỉ có core được log. Các phần khác không được log trừ trường hợp đặc biệt là chạy config
* Các thành phần khác ko đươc hiểu nhau

* Core không cân hiểu các phần khác nhau của hệ thống mà hiểu các ánh xạ.

* Core luôn có các kết nối tới các phần bằng drive

* Logic ko được để ngoài core trừ trường hợp đặc hữu dành riêng cho drive Vd: bson.NewObjectID() là thành phần của mongo

## Sơ đồ hệ thống

```text
                  +-------+
                  |       |
                  | REDIS +------+                   +-------------+
                  |       |      |                   |             |
                  +-------+      |                   |             |
                                (2)                  |             |
+-----------+                    |                   |             +-------(6)--+
|           |                    |                   |             |            |
| CONFIG ENV|                   +v-------+           |             |        +---v-----+
|           |                   |        |           |             |        |         |        ^
+--------------------(0)-------->  MAIN  +----(4)---->             |        |   GIN   +--------+
      +-------------------(3)--->        |           |    CORE     |        |         |
      |     |                   ++-------+           |             |        +---+-----+
      | WS  |                    ^                   |             |            |
      |     |                    |                   |             <-----(5)----+
      +-----+                    |                   |             |
                                (1)                  |             |
                  +--------+     |                   |             |
                  |        |     |                   |             |
                  | MONGO  +-----+                   |             |
                  |        |                         +-------------+
                  +--------+

```
 * (0) Load config
 * (1),(2),(3) Initial adapter mongo, redis, websoket
 * (4) Initial Core, start core, add agument for core
 * (5) Core receive process from gin
 * (6) Core process then return result to gin
### Setup
 * Online here: [social](https://serene-headland-81432.herokuapp.com/)
#### 1. Knowleage
 * Golang office docs here [The Go programming language](https://golang.org/doc/)
 * Golang by example [Go example](https://gobyexample.com/)
 * Mongo office docs here [MongoDb](https://docs.mongodb.com/?_ga=2.47569767.1996944194.1521029475-988272004.1521029475)
 * Redis office docs here [Redis](https://redis.io/documentation)

#### 2. Get project

```bash
git clone https://github.com/my0sot1s/social.git
cd social 
dep ensure -update
go run main.go
# app run at localhost:4444
```
 * default `$PORT` =  `4444`
 * local can run with my domain : `http://local.tenm.cf:4444`


### Core code 

#### 1. Core Interface
```go
type ICore interface {
	GetUserByUname(username string) (error, *m.User)
	GetUserByEmail(email string) (error, *m.User)
	GetPost(limit, page int, userID string) (error, []*m.Post)
	GetPostById(postID string) (error, *m.Post)
	GetFeed(limit, page int, userID string) (error, []*m.Feed)
	GetFollower(own string) (error, []*m.Follower)
	GetFollowing(follower string) (error, []*m.Follower)
	CountLike(postID string) (error, int)
	GetAlbum(AlbumID string) (error, *m.Album)
	GetComments(limit, page int, postID string) (error, []*m.Comment)
	GetLikes(postID string) (error, []*m.Like)
	IsUserLikePost(pid, uid string) (error, bool)
	GetPosts(pIDs []string) (error, []*m.Post)
	GetUserOwns(uIDs []string) (error, []*m.User)
	CreatePost(p *m.Post) (error, *m.Post)
	CreateComment(c *m.Comment) (error, *m.Comment)
	CreateFeed(f *m.Feed) (error, *m.Feed)
	CreateFeeds(feeds []*m.Feed) (error, []interface{})
	CreateUser(u *m.User) (error, *m.User)
	ModifyFollower(t *m.Follower) (error, *m.Follower)
	CreateAlbum(a *m.Album) (error, *m.Album)
	HitLikePost(postID, userID string) error
	UnlikePost(postID, userID string) error
	//
	GetUsersLikePost(userIDs []string) (error, []*m.User)
	FollowUser(f *m.Follower) error
	UnfollowUser(own, uid string) error
}
```

#### 2. API mapping

```go 
// g.router.GET("/favicon.ico", g.sendFavicon)
	// follow
	follow := g.router.Group("/follow")
	follow.GET("/:uid/follower", g.getListFollower)
	follow.GET("/:uid/following", g.getListFollowing)
	follow.POST("do/:uid/follow", g.followAnUser)
	follow.POST("do/:uid/unfollow", g.unFollowAnUser)

	// like
	like := g.router.Group("/like")
	like.GET("pid/:pid", g.countLikeAPost)
	like.POST("/:pid/like", g.hitLikeAPost)
	like.POST("/:pid/unlike", g.unlikeAPost)
	like.GET("users/:pid", g.getUserLikeAPost)
	// like.PUT("/:id/byUser", g.checkOwnerLikePost)

	// feed
	feed := g.router.Group("/feed")
	feed.GET("uid/:uid", g.getUserFeed)
	feed.GET("posts/:uid", g.getFeedPostByUid)

	// post
	post := g.router.Group("/post")
	post.GET("uid/:uid", g.getUserPost)
	post.GET("pid/:pid", g.getUserPostByID)

	//comment
	comment := g.router.Group("/comment")
	comment.GET("pid/:pid", g.getCommentPost)
	comment.POST("pid/:pid", g.addCommentToPost)

	// auth
	g.router.POST("login", g.Login)
	g.router.POST("register", g.Register)
```

#### Các link tham khảo về thiết kế hệ thống.
 * Cách sử dụng 1 số trick nhỏ khi làm js [https://jrsinclair.com/articles/2016/gentle-introduction-to-functional-javascript-functions/](https://jrsinclair.com/articles/2016/gentle-introduction-to-functional-javascript-functions/)
 * Kiến trúc sạch - clean architecture là gì? [Kiến trúc sạch - clean architecture](https://techmaster.vn/posts/34271/kien-truc-sach-clean-architecture)
 * A proposal for bundling reducers, action types and actions when using Redux [Ducks: Redux Reducer Bundles](https://github.com/erikras/ducks-modular-redux)
 * `redux-saga` is a library for react [redux-saga](https://github.com/redux-saga/redux-saga)
 * [Redux is a predictable state container for JavaScript apps.](https://redux.js.org/)