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
