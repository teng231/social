package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) getAlbumById(ctx *gin.Context) {
	abId := ctx.Param("abId")
	if abId == "" {
		ctx.JSON(400, gin.H{
			"error": "no abId",
		})
		return
	}
	err, albumInfo := g.cr.LoadAlbumById(abId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"album": albumInfo,
	})
}

func (g *GinConfig) getAlbumByAuthor(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if authorId == "" {
		ctx.JSON(400, gin.H{
			"error": "no authorId",
		})
		return
	}
	err, albums := g.cr.LoadAlbumByAuthor(limit, page, authorId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"albums": albums,
	})
}

func (g *GinConfig) createUserAlbum(ctx *gin.Context) {
	authorId := ctx.Param("uid")
	albumName := ctx.PostForm("name")
	ownerID := ctx.PostForm("owner")
	media := ctx.PostForm("media")
	if authorId == "" {
		ctx.JSON(400, gin.H{
			"error": "no authorId",
		})
		return
	}
	if ownerID == "" {
		ctx.JSON(400, gin.H{
			"error": "no owner",
		})
		return
	}
	if media == "" {
		ctx.JSON(400, gin.H{
			"error": "no media",
		})
		return
	}
	err, album := g.cr.UpsertAnAlbum(albumName, media, ownerID)
	if err == nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"albums": album,
	})
}
