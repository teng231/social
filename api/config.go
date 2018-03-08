package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinConFig struct {
	router *gin.Engine
	PORT   int
	mode   string
}

func (g *GinConFig) Config(port int, mode string) {
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	g.router = gin.New()
	g.router.Use(gin.Recovery())
	g.router.Use(gin.Logger())
	g.PORT = port
}

func (g *GinConFig) Run() {
	g.GinStart()
	g.router.Run(fmt.Sprintf(":%d", g.PORT))
}

func (g *GinConFig) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"pong": "ok",
	})
}
func (g *GinConFig) GinStarted(ctx *gin.Context) {
	ctx.String(200, "Gin started")
}

func (g *GinConFig) GinStart() {
	g.router.GET("ping", g.Ping)
	g.router.GET("", g.GinStarted)
}
