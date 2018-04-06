package api

import (
	"github.com/gin-gonic/gin"
)

func (g *GinConfig) createNewEmotion(ctx *gin.Context) {
	uid := ctx.Param("uid")
	by := ctx.PostForm("by")
	medias := ctx.PostForm("medias")
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	if by == "" {
		ctx.JSON(400, gin.H{
			"error": "no by",
		})
		return
	}
	if medias == "" {
		ctx.JSON(400, gin.H{
			"error": "medias",
		})
		return
	}
	m := g.cr.CreateEmotion(medias, by)
	ctx.JSON(200, gin.H{
		"emotion": m,
	})
}

func (g *GinConfig) getAllEmotions(ctx *gin.Context) {
	uid := ctx.Param("uid")
}
