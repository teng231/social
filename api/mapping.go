package api

func (g *GinConfig) ginStart() {
	g.router.GET("ping", g.ping)
	g.router.GET("", g.ginStarted)

	g.router.GET("feed/:uid/:limit/:page", g.getUserFeed)
	g.router.GET("post/:uid/:limit/:page", g.getUserPost)
	g.router.GET("posts/:pid", g.getUserPostByID)

	g.router.GET("comments/:pid/:limit/:page", g.getCommentPost)
}
