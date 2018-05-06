package api

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) ReadSingleUserInfo(ctx *gin.Context) {
	uid := ctx.Param("uid")
	err, u := g.cr.GetUserInfo(uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"user": u,
	})
}
func (g *GinConfig) ReadMultipleUserInfo(ctx *gin.Context) {
	// uid := ctx.Param("uid")
	uids := ctx.PostForm("uIDs")
	listUserIds := make([]string, 0)
	err := json.Unmarshal([]byte(uids), &listUserIds)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	err, u := g.cr.GetMultipleUserInfo(listUserIds)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"users": u,
	})
}
func (g *GinConfig) SearchForUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	query := ctx.Query("query")
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "uid not found",
		})
		return
	}
	if &query == nil {
		query = `.*`
	}
	err, us := g.cr.LookupUserByQuery(query)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"users": us,
	})
}
