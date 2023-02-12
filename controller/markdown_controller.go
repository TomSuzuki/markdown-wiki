package controller

import (
	"net/http"
	"strings"

	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// MarkdownController ...Markdownとして単語を取得する。
func MarkdownController(c *gin.Context) {
	// query[word]
	word, _ := c.GetQuery("word")

	// get markdwon
	md, _ := model.GetWordText(word)
	wordFolder := strings.Split(word, "/")
	name := wordFolder[len(wordFolder)-1]

	// send
	c.JSON(http.StatusOK, view.MarkdDownData{
		Word:     name,
		Markdown: md,
		FileName: name + ".md",
	})
}
