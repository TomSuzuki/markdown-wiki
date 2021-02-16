package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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

	// path
	path := fmt.Sprintf("%s%s.md", config.PageSavePath, name)

	// path ok?
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)

	// save
	ioutil.WriteFile(path, []byte(text), 0666)
	c.String(http.StatusOK, "")
}
