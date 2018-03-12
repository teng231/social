package api

import (
	"fmt"
	"strconv"

	"github.com/my0sot1s/social/core"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_LIMIT = 20
	DEFAULT_PAGE  = 1
)

type GinConfig struct {
	router *gin.Engine
	PORT   string
	mode   string
	cr     *core.Core
}

// Config is a constructer
func (g *GinConfig) Config(port, mode string, cr *core.Core) {
	if mode == "" {
		mode = gin.ReleaseMode
	}
	// set mode `production` or `dev`
	gin.SetMode(mode)
	g.router = gin.New()
	g.router.Use(gin.Recovery())
	g.router.Use(gin.Logger())
	g.PORT = port
	g.cr = cr
	g.router.StaticFile("/favicon.ico", "./../favicon.ico")
}

// Run start api
func (g *GinConfig) Run() {
	g.ginStart()
	g.router.Run(fmt.Sprintf(":%v", g.PORT))
}

func (g *GinConfig) ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"pong": "ok",
	})
}

func (g *GinConfig) ginStarted(ctx *gin.Context) {
	ctx.String(200, "Gin started")
}

func (g *GinConfig) getLimitPage(strLimit, strPage string) (int, int) {
	limit, err1 := strconv.Atoi(strLimit)
	page, err2 := strconv.Atoi(strPage)

	if err1 != nil {
		limit = DEFAULT_LIMIT
	}

	if err2 != nil {
		page = DEFAULT_PAGE
	}

	return limit, page
}
