package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TomSuzuki/markdown-wiki/config"
)

// RemoveDirectory ...再帰的にフォルダを削除する。
func RemoveDirectory(dir string) {
	dirlist := strings.Split(filepath.ToSlash(dir), "/")
	temp := dir
	for i := len(dirlist) - 1; i >= 0; i-- {
		temp = temp[:strings.LastIndex(temp, dirlist[i])]
		os.Remove(temp)
	}
}

// RemoveWord ...単語を指定してファイルを削除する。
func RemoveWord(word string) error {
	return os.Remove(fmt.Sprintf("%s%s.md", config.PageSavePath, word))
}
