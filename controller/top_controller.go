package controller

import (
	"fmt"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// TopPageController ...トップページを表示する（内部は単語ページ流用する）。
func TopPageController(c *gin.Context) {
	// dto
	var data view.WordPage
	data.Title = config.ServiceName
	data.MarkdownText, _ = model.GetFileString(config.TopPageMarkdownPath)
	data.MarkdownHTML = model.MarkdownToHTML(data.MarkdownText)
	data.CanEdit = false

	// view
	view.NewView(c, view.PageData{
		PageTitle: fmt.Sprintf("%s | %s", config.ServiceName, "Top"),
		HTML:      view.WordPageView(data),
		MenuInfo: view.MenuInfo{
			MenuTop: true,
		},
	})
}
