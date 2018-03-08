package api

import (
	"github.com/gin-gonic/gin"
)

func (g *GinConfig) getUserFeed(ctx *gin.Context) {
	uid := ctx.Param("uid")
	limit, page := g.getLimitPage(ctx.Param("limit"), ctx.Param("page"))

	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
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
	limit, page := g.getLimitPage(ctx.Param("limit"), ctx.Param("page"))

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
	limit, page := g.getLimitPage(ctx.Param("limit"), ctx.Param("page"))
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
