package model

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TomSuzuki/markdown-wiki/config"
)

// GetMarkdownText ...マークダウンを読み取って返す。
func GetMarkdownText(path string) (string, error) {
	os.Mkdir(config.PageSavePath, 0777)
	md, err := ioutil.ReadFile(fmt.Sprintf("%s%s.md", config.PageSavePath, path))
	return string(md), err
}
