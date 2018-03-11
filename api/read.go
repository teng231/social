package api

import (
	"github.com/gin-gonic/gin"
)

func (g *GinConfig) getUserFeed(ctx *gin.Context) {
	uid := ctx.Param("uid")
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no user id",
		})
	}
	err, feeds := g.cr.GetFeedByUser(limit, page, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, feeds)
}

func (g *GinConfig) getUserPost(ctx *gin.Context) {
	uid := ctx.Param("uid")
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}

	err, posts := g.cr.GetPostUser(limit, page, uid)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, posts)
}

func (g *GinConfig) getUserPostByID(ctx *gin.Context) {
	pid := ctx.Param("pid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}
	err, post := g.cr.GetPostID(pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, post)
}

func (g *GinConfig) getCommentPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}
	err, comments := g.cr.GetCommentByPostID(limit, page, pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, comments)
}

func (g *GinConfig) getFeedPostByUid(ctx *gin.Context) {
	userID := ctx.Param("uid")
	// get feed
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if userID == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
	}
	// get post
	err, posts, users := g.cr.GetPostsByFeedUser(limit, page, userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"posts": posts,
		"users": users,
	})
}

func (g *GinConfig) addCommentToPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	text := ctx.PostForm("text")
	userID := ctx.PostForm("user_id")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}
	err, comments := g.cr.InsertCommentsToPost(pid, text, userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, comments)
}

func (g *GinConfig) countLikeAPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}
	err, count := g.cr.GetCountLike(pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, count)
}

func (g *GinConfig) unlikeAPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	uid := ctx.PostForm("uid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
	}
	err := g.cr.UnlikePost(pid, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{"unlike": pid})
}

func (g *GinConfig) hitLikeAPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	uid := ctx.PostForm("uid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
	}
	err := g.cr.HitLikePost(pid, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{"like": pid})
}

func (g *GinConfig) getUserLikeAPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
	}
	err, users := g.cr.GetUserLikePost(pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{"users": users})
}
