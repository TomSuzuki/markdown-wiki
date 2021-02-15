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
	router.GET("/error", controller.ErrorPageController)
	router.GET("/page", controller.WordPageController)

	// no route
	// rootPath.GET("", ctrl.Top)
	router.NoRoute(controller.ErrorPageController)

	return
}
