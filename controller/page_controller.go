package controller

import (
	"html/template"

	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// WordPageController ...単語のページを表示します。
func WordPageController(c *gin.Context) {
	// query
	word, err := service.QueryString(c, "w")
	if err != nil {
		view.NewView(c, view.PageData{
			HTML: view.ErrorPageView(view.ErrorPage{
				ErrorCode: "030",
			}),
			MenuInfo: view.MenuInfo{
				MenuSearch: true,
			},
		})
		return
	}

	// dto
	var data view.WordPage
	data.Word = word
	data.MarkdownText, err = model.GetMarkdownText(word)
	data.MarkdownHTML = template.HTML(string(blackfriday.MarkdownCommon([]byte(data.MarkdownText))))
	if err != nil {
		// 新規作成ページに飛ばす。
	}

	// view
	view.NewView(c, view.PageData{
		HTML: view.WordPageView(data),
		MenuInfo: view.MenuInfo{
			MenuSearch: true,
			Word:       word,
		},
	})
}
