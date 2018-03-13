package api

func (g *GinConfig) ginStart() {
	// ping pong
	g.router.GET("ping", g.ping)
	g.router.GET("", g.ginStarted)
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
}
