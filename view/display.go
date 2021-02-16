package view

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/gin-gonic/gin"
)

// NewView ...表示する。
func NewView(c *gin.Context, data PageData) {
	if data.PageTitle == "" {
		data.PageTitle = config.ServiceName
	}
	c.HTML(http.StatusOK, "index.html", data)
}

// createHTML ...htmlの生成。
func createHTML(file string, data interface{}) template.HTML {
	var body bytes.Buffer
	t := template.Must(template.ParseFiles(file))
	t.Execute(&body, data)
	return template.HTML(body.String())
}

// ErrorPageView ...エラーページ部分。
func ErrorPageView(data ErrorPage) template.HTML {
	data.ErrorMessage = errorMessage[data.ErrorCode]
	return createHTML("templates/error.html", data)
}

// WordPageView ...単語のページの表示。
func WordPageView(data WordPage) template.HTML {
	return createHTML("templates/word.html", data)
}

// WritePageView ...新規作成ページの表示。
func WritePageView(data WritePage) template.HTML {
	return createHTML("templates/write.html", data)
}

// SearchPageView ...検索ページの表示。
func SearchPageView(data SearchPage) template.HTML {
	return createHTML("templates/search.html", data)
}

// EditPageView ...編集ページの表示（内部は新規作成と同じ）。
func EditPageView(data EditPage) template.HTML {
	return createHTML("templates/write.html", data)
}
