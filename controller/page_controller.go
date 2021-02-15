package controller

import (
	"html/template"

	"github.com/Depado/bfchroma"
	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
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

	// rendering setting
	r := bfchroma.NewRenderer(
		bfchroma.WithoutAutodetect(),
		bfchroma.ChromaStyle(styles.Dracula),
		bfchroma.ChromaOptions(html.WithLineNumbers(false)),
		bfchroma.Extend(
			blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
				Flags: blackfriday.UseXHTML | blackfriday.Smartypants | blackfriday.SmartypantsFractions |
					blackfriday.SmartypantsDashes | blackfriday.SmartypantsLatexDashes | blackfriday.TOC,
			}),
		),
	)

	// dto
	var data view.WordPage
	data.Word = word
	data.MarkdownText, err = model.GetMarkdownText(word)
	//data.MarkdownHTML = template.HTML(string(blackfriday.Run([]byte(data.MarkdownText))))
	data.MarkdownHTML = template.HTML(blackfriday.Run([]byte(data.MarkdownText), blackfriday.WithRenderer(r)))
	if err != nil {
		// 新規作成ページに飛ばす。
	}

	// view
	view.NewView(c, view.PageData{
		HTML: view.WordPageView(data),
		MenuInfo: view.MenuInfo{
			MenuSearch: true,
		},
	})
}
