package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/mirrors"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "" || password == "" {
		ctx.JSON(400, gin.H{
			"error": "username or password nil",
		})
		return
	}
	err, user, token := g.cr.Login(username, password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}

func (g *GinConfig) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	if username == "" || password == "" || email == "" {
		ctx.JSON(400, gin.H{
			"error": "username or password or email nil",
		})
		return
	}
	err, user := g.cr.Register(&mirror.User{
		UserName: username,
		Password: password,
		Email:    email,
	})
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

func (g *GinConfig) confirmToken(ctx *gin.Context) {
	token := ctx.Param("token")
	uid := ctx.Param("uid")
	err := g.cr.CheckKeyToken(token, uid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"state": "activated",
	})

}
