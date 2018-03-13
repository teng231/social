package api

func (g *GinConfig) ginStart() {
	// ping pong
	g.router.GET("ping", g.ping)
	g.router.GET("", g.ginStarted)
	// g.router.GET("/favicon.ico", g.sendFavicon)
	// album
	album := g.router.Group("/album")
	album.GET("/:uid/byId/:abId", g.getAlbumById)
	album.GET("/:uid/byAuthor/:authorId", g.getAlbumByAuthor)
	// follow
	follow := g.router.Group("/follow")
	follow.GET("/:uid/follower/:userTarget", g.getListFollower)
	follow.GET("/:uid/following/:userTarget", g.getListFollowing)
	follow.POST("/:uid/follow", g.followAnUser)
	follow.POST("/:uid/unfollow", g.unFollowAnUser)

	// like
	like := g.router.Group("/like")
	like.GET("/:uid/count/:pid", g.countLikeAPost)
	like.GET("/:uid/users/:pid", g.getUserLikeAPost)
	like.POST("/:uid/like/:pid", g.hitLikeAPost)
	like.POST("/:uid/unlike/:pid", g.unlikeAPost)
	like.POST("/:uid/owner/:pid", g.checkOwnerLikePost)

	// feed
	feed := g.router.Group("/feed")
	feed.GET("/:uid/userFeed/:userTarget", g.getUserFeed)
	feed.GET("/:uid/feedPost/:userTarget", g.getFeedPostByUid)

	// post
	post := g.router.Group("/post")
	post.GET("/:uid", g.getUserPost)
	post.GET("/:uid/post/:pid", g.getPostByID)

	//comment
	comment := g.router.Group("/comment")
	comment.GET("/:uid/post/:pid", g.getCommentPost)
	comment.POST("/:uid/post/:pid", g.addCommentToPost)

	// auth
	g.router.POST("login", g.Login)
	g.router.POST("register", g.Register)
}
