package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/gin-gonic/gin"
)

// SaveController ...データの保存を行う。
func SaveController(c *gin.Context) {

	// form
	name, b := c.GetPostForm("name")
	if !b || name == "" {
		c.String(http.StatusBadRequest, "")
		return
	}
	text, b := c.GetPostForm("text")
	if !b {
		c.String(http.StatusBadRequest, "")
		return
	}

	// save
	path := fmt.Sprintf("%s%s.md", config.PageSavePath, name)
	ioutil.WriteFile(path, []byte(text), 0666)
	c.String(http.StatusOK, "")
}
