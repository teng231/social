package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) getUserFeed(ctx *gin.Context) {
	userTarget := ctx.Param("userTarget")
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if userTarget == "" {
		ctx.JSON(400, gin.H{
			"error": "no userTarget id",
		})
		return
	}
	err, feeds := g.cr.LoadFeedByUser(limit, page, userTarget)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"feeds": feeds})
}
