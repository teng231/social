package api

import "github.com/gin-gonic/gin"

func (g *GinConfig) ginStart() {
	// ping pong
	g.router.GET("ping", g.ping)
	// g.router.GET("", g.ginStarted)
	g.router.GET("signature-file", g.signatureFileToUpload)
	g.router.POST("postDemo/:uid", g.postDemo)
	g.router.LoadHTMLGlob("api/upload/*")
	g.router.LoadHTMLFiles("api/upload/index.html")
	g.router.GET("upload", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		ctx.HTML(200, "index.html", gin.H{})
	})
	// g.router.GET("/favicon.ico", g.sendFavicon)
	// emotion
	emotion := g.router.Group("/emotion")
	emotion.GET("/:uid", g.getAllEmotions)
	emotion.POST("/:uid/create", g.createNewEmotion)
	// saved
	saved := g.router.Group("/saved")
	saved.POST("/:uid/create", g.SavePost)
	saved.GET("/:uid", g.getSaved)
	// album
	album := g.router.Group("/album")
	album.GET("/:uid/byId/:abId", g.getAlbumById)
	album.GET("/:uid/byAuthor/:authorId", g.getAlbumByAuthor)
	album.POST("/:uid/create", g.createUserAlbum)
	// follow
	follow := g.router.Group("/follow")
	follow.GET("/:uid/follower/:userTarget", g.getListFollower)
	follow.GET("/:uid/following/:userTarget", g.getListFollowing)
	follow.GET("/:uid/count/:userTarget", g.getCountFollow)
	follow.POST("/:uid/unfollow", g.unFollowAnUser)
	follow.POST("/:uid/follow", g.followAnUser)
	follow.GET("/:uid/is/:userTarget", g.checkFollowOtherUser)

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
	post.POST("/:uid/create", g.createNewPost)
	post.DELETE("/:uid/delete", g.createNewPost)

	//comment
	comment := g.router.Group("/comment")
	comment.GET("/:uid/post/:pid", g.getCommentPost)
	comment.GET("/:uid/count/:pid", g.countCommentByPost)
	comment.POST("/:uid/post/:pid", g.addCommentToPost)

	// auth
	g.router.POST("login", g.Login)
	g.router.POST("register", g.Register)
	g.router.GET("confirm/:uid/:token", g.confirmToken)
	// g.router.GET("/abc", func(ctx *gin.Context) {
	// 	err := errors.New("1111")
	// 	ctx.JSON(200, gin.H{
	// 		"error": utils.ErrStr(err),
	// 	})
	// })
	user := g.router.Group("/user")
	user.GET("/:uid", g.ReadSingleUserInfo)
	user.POST("/multiples", g.ReadMultipleUserInfo)
	user.GET("/:uid/search", g.SearchForUser)

	explore := g.router.Group("/explore")
	explore.GET("/:uid", g.getExplore)

}
