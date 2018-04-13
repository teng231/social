package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) getExplore(ctx *gin.Context) {
	uid := ctx.Param("uid")
	limit, anchor := g.getLimitPage(ctx.Query("limit"), ctx.Query("anchor"))
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	err, ps, anchor := g.cr.GetAnyPost(limit, anchor, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"explores": ps,
		"anchor":   anchor,
	})
}
