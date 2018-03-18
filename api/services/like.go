package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) countLikeAPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}
	err, count := g.cr.LoadCountLike(pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
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
		return
	}
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	err := g.cr.RemoveLikePost(pid, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
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
	err := g.cr.UpsertLikePost(pid, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
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
		return
	}
	err, users := g.cr.LoadUserLikePost(pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{"users": users})
}

func (g *GinConfig) checkOwnerLikePost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	uid := ctx.PostForm("uid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	err, ok := g.cr.CheckOwnerLikePost(pid, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"uid":    uid,
		"pid":    pid,
		"status": ok,
	})
}
