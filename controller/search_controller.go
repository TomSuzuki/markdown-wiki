package controller

import (
	"net/url"
	"path/filepath"
	"strings"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/model"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// SearchPageController ...検索ページの処理。
func SearchPageController(c *gin.Context) {
	// query[keyword]
	keyword, _ := c.GetQuery("keyword")

	// query[path]
	path, _ := c.GetQuery("f")
	path, _ = url.QueryUnescape(path)
	path = model.DirPathClean(path)

	// data
	var data view.SearchPage
	data.Keyword = keyword
	data.Path = path

	// keyword mode
	if keyword != "" {
		// get list
		list := model.Dirwalk(config.PageSavePath, false)

		// filter -> append
		for i := range list {
			list[i] = model.GetFileNameWithoutExt(list[i])
			if strings.Contains(list[i], keyword) {
				data.WordList = append(data.WordList, view.PathData{
					Path:     list[i],
					PathName: list[i],
				})
			}
		}
	} else {
		// get list
		data.FolderList = model.GetReadDir(path, true)
		data.WordList = model.GetReadDir(path, false)

		// without extention
		for i := range data.WordList {
			data.WordList[i].Path = model.GetFileNameWithoutExt(data.WordList[i].Path)
			data.WordList[i].PathName = model.GetFileNameWithoutExt(data.WordList[i].PathName)
		}

		// add[..]
		if model.IsRoot(path) {
			data.FolderList = append([]view.PathData{{
				Path:     filepath.Dir(path[:len(path)-2]),
				PathName: "..",
			}}, data.FolderList...)
		}
	}

	// view
	view.NewView(c, view.PageData{
		HTML: view.SearchPageView(data),
		MenuInfo: view.MenuInfo{
			MenuSearch: true,
		},
	})
}
