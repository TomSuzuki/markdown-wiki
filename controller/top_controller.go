package controller

import (
	"html/template"
	"io/ioutil"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// TopPageController ...トップページを表示する（内部は単語ページ流用する）。
func TopPageController(c *gin.Context) {
	// md
	md, _ := ioutil.ReadFile(config.TopPageMarkdownPath)

	// rendering
	extensionsFlags := blackfriday.EXTENSION_FENCED_CODE
	htmlFlags := blackfriday.HTML_TOC
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	options := blackfriday.Options{Extensions: extensionsFlags}

	// dto
	var data view.WordPage
	data.Word = config.ServiceName
	data.MarkdownText = string(md)
	data.MarkdownHTML = template.HTML(string(blackfriday.MarkdownOptions([]byte(data.MarkdownText), renderer, options)))

	// view
	view.NewView(c, view.PageData{
		HTML: view.WordPageView(data),
		MenuInfo: view.MenuInfo{
			MenuTop: true,
		},
	})
}
