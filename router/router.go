package router

import (
	"HideSeekCatGo/handler/sd"
	"HideSeekCatGo/handler/user"
	"HideSeekCatGo/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
	}

	// the health check handlers.
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
