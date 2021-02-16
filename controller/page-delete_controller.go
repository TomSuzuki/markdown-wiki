package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/gin-gonic/gin"
)

// DeletePageController ...指定のページを削除する。
func DeletePageController(c *gin.Context) {
	// query
	word, err := service.QueryString(c, "w")
	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}

	// delete
	path := fmt.Sprintf("%s%s.md", config.PageSavePath, word)
	if err := os.Remove(path); err != nil {
		c.String(http.StatusNotFound, "")
		return
	}

	// return
	c.String(http.StatusOK, "")
}
