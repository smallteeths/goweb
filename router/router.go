package router

import (
	"net/http"

	"tool-backend/service"
	"tool-backend/router/middleware"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化路由
func InitRouter(g *gin.Engine) {
	middlewares := []gin.HandlerFunc{}
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	router := g.Group("/upload")
	{
		router.POST("/logo", service.UploadHandler)

		router.POST("/loginBackground", service.UploadBackgroundHandler)

		router.POST("/save", service.Save)

		router.POST("/variableLastConfig", service.SelectTemplateVariable)

		router.GET("/test", service.Test)

		// router.StaticFS("/public", http.Dir("static/uploadfile/"))

		router.StaticFS("/public", http.Dir("./"))

	}

}