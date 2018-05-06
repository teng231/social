package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my0sot1s/social/utils"
)

func (g *GinConfig) createNewPost(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userID := ctx.PostForm("user_id")
	content := ctx.PostForm("content")
	tags := ctx.PostForm("tags")
	medias := ctx.PostForm("medias")
	if uid == "" {
		utils.Log("no uid")
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	if userID == "" {
		utils.Log("no userID")

		ctx.JSON(400, gin.H{
			"error": "no userID",
		})
		return
	}
	if content == "" {
		utils.Log("no no content")
		ctx.JSON(400, gin.H{
			"error": "no content ",
		})
		return
	}
	if medias == "" {
		utils.Log("no no medias")

		ctx.JSON(400, gin.H{
			"error": "no medias",
		})
		return
	}
	if tags == "" {
		utils.Log("no no tags")

		ctx.JSON(400, gin.H{
			"error": "no tags",
		})
		return
	}

	err, post := g.cr.AddNewPostBonusFeed(userID, content, medias, tags)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func (g *GinConfig) getUserPost(ctx *gin.Context) {
	uid := ctx.Param("uid")
	limit, anchor := g.getLimitPage(ctx.Query("limit"), ctx.Query("anchor"))
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}

	err, posts, anchor := g.cr.LoadPostUser(limit, anchor, uid)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"posts":  posts,
		"anchor": anchor,
	})
}

func (g *GinConfig) getPostByID(ctx *gin.Context) {
	pid := ctx.Param("pid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}
	err, post := g.cr.LoadPostID(pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func (g *GinConfig) getFeedPostByUid(ctx *gin.Context) {
	userTarget := ctx.Param("userTarget")
	// get feed
	limit, anchor := g.getLimitPage(ctx.Query("limit"), ctx.Query("anchor"))
	if userTarget == "" {
		ctx.JSON(400, gin.H{
			"error": "no userTarget",
		})
		return
	}
	// get post
	err, posts, anchor := g.cr.LoadPostsByFeedUser(limit, anchor, userTarget)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": utils.ErrStr(err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"posts":  posts,
		"anchor": anchor,
	})
}

func (g *GinConfig) postDemo(ctx *gin.Context) {
	uid := ctx.Param("uid")
	postForm := ctx.PostForm("postForm")
	if postForm == "" {
		ctx.JSON(400, gin.H{
			"error": "postForm is nil",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"success": gin.H{
			"postForm": postForm,
			"uid":      uid,
		},
	})
}
