package controller

import (
	"fmt"
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

	if keyword != "" {
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
		for i := range list {
			data.WordList = append(data.WordList, view.PathData{
				Path:     list[i],
				PathName: list[i],
			})
		}

		// view
		view.NewView(c, view.PageData{
			HTML: view.SearchPageView(data),
			MenuInfo: view.MenuInfo{
				MenuSearch: true,
			},
		})

		return
	}

	// data
	path, _ := c.GetQuery("f")
	var data view.SearchPage
	data.Keyword = keyword
	data.Path = path

	// get list
	dir := fmt.Sprintf("%s/%s", config.PageSavePath, path)
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			data.FolderList = append(data.FolderList, view.PathData{
				Path:     fmt.Sprintf("%s/%s", path, file.Name()),
				PathName: fmt.Sprintf("📁  %s", file.Name()),
			})
		} else {
			data.WordList = append(data.WordList, view.PathData{
				Path:     fmt.Sprintf("%s/%s", path, file.Name()),
				PathName: file.Name(),
			})
		}
	}

	// without extention
	for i := range data.WordList {
		data.WordList[i].PathName = getFileNameWithoutExt(data.WordList[i].PathName)
		data.WordList[i].Path = strings.Replace(getFileNameWithoutExt(data.WordList[i].Path), "/", "", 1)
	}

	// ..
	if path != "." && path != "" && path != "\\" {
		path = filepath.Clean(path)
		path = strings.Replace(path, "\\\\", "", 1)
		path = strings.Replace(path, "//", "", 1)
		data.FolderList = append([]view.PathData{{
			Path:     filepath.Dir(path),
			PathName: "..",
		}}, data.FolderList...)
	}

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
