package server

import "github.com/gin-gonic/gin"

// Router ...ルーティングを行います。
func Router() (router *gin.Engine) {

	// root
	router = gin.Default()
	//router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "./assets")

	// route
	//router.GET("/ping", ctrl.Ping)

	// no route
	// rootPath.GET("", ctrl.Top)
	// router.NoRoute(ctrl.Error)

	return
}
