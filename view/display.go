package view

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewView ...表示する。
func NewView(c *gin.Context, data PageData) {
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
