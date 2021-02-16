package controller

import (
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// WritePageController ...新規作成ページの処理。
func WritePageController(c *gin.Context) {

	// query
	word, _ := service.QueryString(c, "w")

	// dto
	var data view.WritePage
	data.Word = word

	// view
	view.NewView(c, view.PageData{
		HTML: view.WritePageView(data),
		MenuInfo: view.MenuInfo{
			MenuNewPage: true,
		},
	})
}
