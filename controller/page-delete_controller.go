package controller

import (
	"fmt"
	"net/http"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/gin-gonic/gin"
)

// DeletePageController ...指定のページを削除する。
func DeletePageController(c *gin.Context) {
	// query[word]
	word, err := service.QueryString(c, "w")
	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}

	// delete
	if err := model.RemoveWord(word); err != nil {
		c.String(http.StatusNotFound, "")
		return
	}

	// empty folder remove
	model.RemoveDirectory(fmt.Sprintf("%s%s.md", config.PageSavePath, word))

	// return
	c.String(http.StatusOK, "")
}
