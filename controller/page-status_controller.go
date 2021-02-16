package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/server/service"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// PageStatusController ...ページの情報を返す。
func PageStatusController(c *gin.Context) {
	// query
	name, err := service.QueryString(c, "name")
	if err != nil || name == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	// data
	var data view.PageStatus
	path := fmt.Sprintf("%s%s.md", config.PageSavePath, name)
	_, err = os.Stat(path)
	data.Exist = err == nil

	// send
	c.JSON(http.StatusOK, data)
}
