package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/mongo"
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
			"error": err,
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
	err, user, token := g.cr.Register(&mongo.User{
		UserName: username,
		Password: password,
		Email:    email,
	})
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}