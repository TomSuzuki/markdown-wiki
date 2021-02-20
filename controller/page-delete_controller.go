package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	// folder empty
	dirs := strings.Split(filepath.ToSlash(path), "/")
	temp := path
	for i := len(dirs) - 1; i >= 0; i-- {
		temp = temp[:strings.LastIndex(temp, dirs[i])]
		os.Remove(temp)
	}

	// return
	c.String(http.StatusOK, "")
}
