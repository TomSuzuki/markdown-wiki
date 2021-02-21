package controller

import (
	"fmt"
	"html/template"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// WordPageController ...単語のページを表示します。
func WordPageController(c *gin.Context) {
	// query
	word, err := service.QueryString(c, "w")
	word, _ = url.QueryUnescape(word)
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

	// word link
	word = filepath.ToSlash(filepath.Clean(word))
	wordFolder := strings.Split(word, "/")
	linkPath := "search?f=" + url.QueryEscape("/")
	linkTitle := ""
	for i := range wordFolder[:len(wordFolder)-1] {
		linkPath += url.QueryEscape(fmt.Sprintf("%s/", wordFolder[i]))
		linkTitle += fmt.Sprintf("<a href='%s'>%s</a>/", linkPath, wordFolder[i])
	}
	linkTitle += wordFolder[len(wordFolder)-1]

	// dto
	var data view.WordPage
	data.Word = word
	data.Title = template.HTML(linkTitle)
	data.MarkdownText, _ = model.GetWordText(word)
	data.MarkdownHTML = model.MarkdownToHTML(data.MarkdownText)
	data.CanEdit = true

	// view
	view.NewView(c, view.PageData{
		PageTitle: fmt.Sprintf("%s | %s", config.ServiceName, word),
		HTML:      view.WordPageView(data),
		MenuInfo: view.MenuInfo{
			MenuSearch: true,
			Word:       word,
		},
	})
}
