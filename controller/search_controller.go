package controller

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/view"
	"github.com/gin-gonic/gin"
)

// SearchPageController ...検索ページの処理。
func SearchPageController(c *gin.Context) {

	// query
	keyword, _ := c.GetQuery("keyword")

	// get list
	list := dirwalk(config.PageSavePath, false)
	for i := range list {
		list[i] = getFileNameWithoutExt(list[i])
	}

	// filter
	listTemp := list
	list = nil
	for i := range listTemp {
		if strings.Contains(listTemp[i], keyword) {
			list = append(list, listTemp[i])
		}
	}

	// data
	var data view.SearchPage
	data.Keyword = keyword
	data.WordList = list

	// view
	view.NewView(c, view.PageData{
		HTML: view.SearchPageView(data),
		MenuInfo: view.MenuInfo{
			MenuSearch: true,
		},
	})
}

// getFileNameWithoutExt ...文字列から拡張子を削除。
func getFileNameWithoutExt(path string) string {
	return strings.Replace(filepath.Join(filepath.Dir(path), filepath.Base(path[:len(path)-len(filepath.Ext(path))])), "\\", "/", -1)
}

// dirwalk ...ディレクトリ内ファイルをを再帰的に検索し返す。
func dirwalk(dir string, isIncludeRoot bool) []string {
	files, _ := ioutil.ReadDir(dir)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths2 := dirwalk(filepath.Join(dir, file.Name()), isIncludeRoot)
			if !isIncludeRoot {
				for f := range paths2 {
					paths2[f] = filepath.Join(file.Name(), paths2[f])
				}
			}
			paths = append(paths, paths2...)
			continue
		}
		if isIncludeRoot {
			paths = append(paths, filepath.Join(dir, file.Name()))
		} else {
			paths = append(paths, file.Name())
		}
	}

	return paths
}
