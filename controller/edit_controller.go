package controller

import (
	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// EditPageController ...編集ページの表示処理を行う。
func EditPageController(c *gin.Context) {
	// query
	word, err := service.QueryString(c, "w")
	if err != nil {
		view.NewView(c, view.PageData{
			HTML: view.ErrorPageView(view.ErrorPage{
				ErrorCode: "031",
			}),
			MenuInfo: view.MenuInfo{
				MenuEdit: true,
			},
		})
		return
	}

	// data
	var data view.EditPage
	data.EditName = word
	data.EditText, err = model.GetMarkdownText(word)
	data.IsNew = false

	// view
	view.NewView(c, view.PageData{
		HTML: view.EditPageView(data),
		MenuInfo: view.MenuInfo{
			MenuEdit: true,
			Word:     word,
		},
	})
}
