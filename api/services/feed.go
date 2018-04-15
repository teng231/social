package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) getUserFeed(ctx *gin.Context) {
	userTarget := ctx.Param("userTarget")
	limit, anchor := g.getLimitPage(ctx.Query("limit"), ctx.Query("anchor"))
	if userTarget == "" {
		ctx.JSON(400, gin.H{
			"error": "no userTarget id",
		})
		return
	}
	err, feeds, newAnchor := g.cr.LoadFeedByUser(limit, anchor, userTarget)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"feeds":  feeds,
		"anchor": newAnchor,
	})
}
