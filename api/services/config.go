package api

import (
	"fmt"
	"os"
	"strconv"

	"github.com/my0sot1s/social/core"
	"github.com/my0sot1s/social/utils"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_LIMIT = 20
)

type GinConfig struct {
	router *gin.Engine
	PORT   string
	mode   string
	cr     *core.Social
}

// Config is a constructer
func (g *GinConfig) Config(port, mode string, cr *core.Social) {
	if mode == "" {
		mode = gin.TestMode
	}
	// set mode `production` or `dev`
	gin.SetMode(mode)
	g.router = gin.New()
	g.router.Use(gin.Recovery())
	g.router.Use(gin.Logger())
	g.PORT = port
	g.cr = cr
	g.router.Use(g.middlewareHeader())
	g.router.StaticFile("/favicon.ico", "./../favicon.ico")
}

// Run start api
func (g *GinConfig) Run() {
	g.ginStart()
	g.router.Run(fmt.Sprintf(":%v", g.PORT))
}

func (g *GinConfig) ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"pong": "sure ok",
	})
}

func (g *GinConfig) ginStarted(ctx *gin.Context) {
	ctx.String(200, "Gin started")
}

func (g *GinConfig) signatureFileToUpload(ctx *gin.Context) {
	signature := g.cr.SignFileToUpload()
	ctx.JSON(200, gin.H{
		"signature": signature,
	})
}

func (g *GinConfig) getLimitPage(strLimit, anchor string) (int, string) {
	limit, err1 := strconv.Atoi(strLimit)
	utils.ErrLog(err1)
	if err1 != nil {
		limit = DEFAULT_LIMIT
	}
	if &anchor == nil {
		anchor = ""
	}
	return limit, anchor
}

func (g *GinConfig) sendFavicon(ctx *gin.Context) {
	ctx.File("../statics/favicon.ico")
}

func (g *GinConfig) middlewareHeader() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			fmt.Println("options")
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.AbortWithStatus(code)
}
func (g *GinConfig) middlewareTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")

		if token == "" {
			respondWithError(401, "API token required", c)
			return
		}

		if token != os.Getenv("API_TOKEN") {
			respondWithError(401, "Invalid API token", c)
			return
		}
		c.Next()
	}
}
