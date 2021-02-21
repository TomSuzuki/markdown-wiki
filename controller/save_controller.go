package controller

import (
	"net/http"

	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/gin-gonic/gin"
)

// SaveController ...データの保存を行う。
func SaveController(c *gin.Context) {
	// form[name]
	word, b := c.GetPostForm("name")
	if !b || word == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	// form[text]
	text, _ := c.GetPostForm("text")

	// save
	model.SaveWord(word, text)

	// return
	c.Status(http.StatusOK)
}
