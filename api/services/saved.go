package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) SavePost(ctx *gin.Context) {
	uid := ctx.Param("uid")
	pid := ctx.PostForm("pid")
	if &pid == nil || pid == "" {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(errors.New("pid empty")),
		})
		return
	}
	err, saved := g.cr.CreateUserSave(uid, pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"saved": saved,
	})
}

func (g *GinConfig) getSaved(ctx *gin.Context) {
	uid := ctx.Param("uid")
	limit, anchor := g.getLimitPage(ctx.Query("limit"), ctx.Query("anchor"))
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no user id",
		})
		return
	}

	err, anchor, saved := g.cr.ListUserSaved(limit, anchor, uid)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"saved":  saved,
		"anchor": anchor,
	})
}
