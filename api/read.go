package api

import (
	"github.com/gin-gonic/gin"
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
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"feeds": feeds})
}

func (g *GinConfig) getUserPost(ctx *gin.Context) {
	uid := ctx.Param("uid")
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}

	err, posts := g.cr.LoadPostUser(limit, page, uid)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"posts": posts,
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
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func (g *GinConfig) getCommentPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}
	err, comments := g.cr.LoadCommentByPostID(limit, page, pid)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"comment": comments,
	})
}

func (g *GinConfig) getFeedPostByUid(ctx *gin.Context) {
	userTarget := ctx.Param("userTarget")
	// get feed
	limit, page := g.getLimitPage(ctx.Query("limit"), ctx.Query("page"))
	if userTarget == "" {
		ctx.JSON(400, gin.H{
			"error": "no userTarget",
		})
		return
	}
	// get post
	err, posts, users := g.cr.LoadPostsByFeedUser(limit, page, userTarget)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"posts": posts,
		"users": users,
	})
}

func (g *GinConfig) addCommentToPost(ctx *gin.Context) {
	pid := ctx.Param("pid")
	text := ctx.PostForm("text")
	userID := ctx.PostForm("uid")
	if pid == "" {
		ctx.JSON(400, gin.H{
			"error": "no post id",
		})
		return
	}
	err, comments := g.cr.UpsertCommentsToPost(pid, text, userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"comment": comments,
	})
}

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
			"error": err,
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
			"error": err,
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
			"error": err,
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
			"error": err,
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
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"uid":    uid,
		"pid":    pid,
		"status": ok,
	})
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
			"error": err,
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
			"error": err,
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
			"error": err,
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
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"ok": true,
	})
}

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
			"error": err,
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
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"albums": albums,
	})
}

func (g *GinConfig) createNewPost(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userID := ctx.PostForm("user_id")
	content := ctx.PostForm("content")
	tags := ctx.PostForm("tags")
	medias := ctx.PostForm("medias")
	if uid == "" {
		ctx.JSON(400, gin.H{
			"error": "no uid",
		})
		return
	}
	if userID == "" || content == "" || medias == "" {
		ctx.JSON(400, gin.H{
			"error": "no userID or content or medias",
		})
		return
	}
	err, albums := g.cr.AddNewPostBonusFeed(userID, content, medias, tags string)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"albums": albums,
	})
}
