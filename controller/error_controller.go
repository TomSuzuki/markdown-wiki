package controller

import (
	"fmt"

	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// ErrorPageController ...エラーページを表示します。
func ErrorPageController(c *gin.Context) {
	// query
	id, _ := service.QueryInt(c, "errorCode")

	// dto
	var data view.ErrorPage
	data.ErrorCode = fmt.Sprintf("%03d", id)

	// view
	view.NewView(c, view.PageData{
		HTML: view.ErrorPageView(data),
		MenuInfo: view.MenuInfo{
			MenuTop: true,
		},
	})
}
