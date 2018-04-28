package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) getCountFollow(ctx *gin.Context) {
	userTarget := ctx.Param("userTarget")
	if userTarget == "" {
		ctx.JSON(400, gin.H{
			"error": "no userTarget",
		})
		return
	}
	err, num := g.cr.CountFollowerByOwner(userTarget)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, num)
}

func (g *GinConfig) getListFollowing(ctx *gin.Context) {
	userTarget := ctx.Param("userTarget")
	if userTarget == "" {
		ctx.JSON(400, gin.H{
			"error": "no userTarget",
		})
		return
	}
	err, users := g.cr.LoadFollowingByUid(userTarget)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"uid":   userTarget,
		"users": users,
	})
}

func (g *GinConfig) getListFollower(ctx *gin.Context) {
	userTarget := ctx.Param("userTarget")
	if userTarget == "" {
		ctx.JSON(400, gin.H{
			"error": "no userTarget",
		})
		return
	}
	err, users := g.cr.LoadFollowerByOwner(userTarget)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"uid":   userTarget,
		"users": users,
	})
}

func (g *GinConfig) followAnUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	owner := ctx.PostForm("owner")
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	if owner == "" {
		ctx.JSON(400, gin.H{
			"error": "no owner",
		})
		return
	}
	err := g.cr.UpsertFollowAnUser(uid, owner)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"ok": true,
	})
}

func (g *GinConfig) unFollowAnUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	owner := ctx.PostForm("owner")
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	if owner == "" {
		ctx.JSON(400, gin.H{
			"error": "no owner",
		})
		return
	}
	err := g.cr.RemoveFollowAnUser(uid, owner)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"ok": true,
	})
}
