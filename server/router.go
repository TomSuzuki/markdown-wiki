package server

import (
	"github.com/TomSuzuki/markdown-wiki/controller"
	"github.com/gin-gonic/gin"
)

// Router ...ルーティングを行います。
func Router() (router *gin.Engine) {

	// root
	router = gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "./assets")

	// route
	router.GET("/page/status", controller.PageStatusController)
	router.GET("/error", controller.ErrorPageController)
	router.GET("/page", controller.WordPageController)
	router.GET("/top", controller.TopPageController)
	router.GET("/write", controller.WritePageController)
	router.GET("/search", controller.SearchPageController)
	router.GET("/edit", controller.EditPageController)
	router.GET("/markdown", controller.MarkdownController)

	router.POST("/save", controller.SaveController)

	router.DELETE("/page", controller.DeletePageController)

	// no route
	router.GET("", controller.TopPageController)
	router.NoRoute(controller.ErrorPageController)

	return
}
