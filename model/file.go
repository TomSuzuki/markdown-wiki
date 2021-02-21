package model

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/TomSuzuki/markdown-wiki/view"
)

// SaveWord ...ファイルを保存する。
func SaveWord(word string, text string) {
	path := fmt.Sprintf("%s%s.md", config.PageSavePath, word)
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)
	ioutil.WriteFile(path, []byte(text), 0666)
}

// GetFileNameWithoutExt ...文字列から拡張子を削除。
func GetFileNameWithoutExt(path string) string {
	return strings.Replace(filepath.Join(filepath.Dir(path), filepath.Base(path[:len(path)-len(filepath.Ext(path))])), "\\", "/", -1)
}

// Dirwalk ...ディレクトリ内ファイルをを再帰的に検索し返す。
func Dirwalk(dir string, isIncludeRoot bool) []string {
	files, _ := ioutil.ReadDir(dir)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths2 := Dirwalk(filepath.Join(dir, file.Name()), isIncludeRoot)
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

// GetReadDir ...指定ディレクトリからディレクトリ(true)かワード(false)を返す。
func GetReadDir(path string, sw bool) []view.PathData {
	var data []view.PathData
	dir := filepath.Join(config.PageSavePath, path)
	files, _ := ioutil.ReadDir(dir)

	// for display
	template := ""
	if sw {
		template = "📁  %s"
	} else {
		template = "%s"
	}

	// list
	for _, file := range files {
		if file.IsDir() == sw {
			data = append(data, view.PathData{
				Path:     url.QueryEscape(fmt.Sprintf("%s/%s", path, file.Name())),
				PathName: fmt.Sprintf(template, file.Name()),
			})
		}
	}
	return data
}

// DirPathClean ...ディレクトリパスをきれいにする。
func DirPathClean(path string) string {
	path = filepath.Clean(path)
	path = filepath.ToSlash(path)
	if string(path[len(path)-1]) != "/" {
		path += "/"
	}
	if string(path[0]) == "/" {
		path = path[1:]
	}
	return path
}

// IsRoot ...指定パスがルートディレクトリかチェックする（判定が雑い）。
func IsRoot(dir string) bool {
	return dir != "." && dir != "" && dir != "\\" && dir != "/" && dir != "./"
}
