package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) addCommentToPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	text := ctx.PostForm("text")
	userID := ctx.PostForm("uid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}
	err, comments := g.cr.UpsertCommentsToPost(pid, text, userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"comment": comments,
	})
}

func (g *GinConfig) getCommentPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	limit, anchor := g.getLimitPage(ctx.Query("limit"), ctx.Query("anchor"))
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}
	err, comments := g.cr.LoadCommentByPostID(limit, anchor, pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"comment": comments,
	})
}
