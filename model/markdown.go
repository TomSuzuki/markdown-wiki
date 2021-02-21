package model

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/TomSuzuki/markdown-wiki/config"
	"github.com/russross/blackfriday"
)

// GetWordText ...単語ファイルのマークダウンを読み取って返す。
func GetWordText(word string) (string, error) {
	os.Mkdir(config.PageSavePath, 0777)
	return GetFileString(fmt.Sprintf("%s%s.md", config.PageSavePath, word))
}

// GetFileString ...ファイルを読み取って[]byteではなくstringで返す。
func GetFileString(path string) (string, error) {
	text, err := ioutil.ReadFile(path)
	return string(text), err
}

// MarkdownToHTML ...マークダウンからHTMLに。
func MarkdownToHTML(text string) template.HTML {
	// option
	extensionsFlags := blackfriday.EXTENSION_FENCED_CODE
	htmlFlags := blackfriday.HTML_TOC
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	options := blackfriday.Options{Extensions: extensionsFlags}

	return template.HTML(string(blackfriday.MarkdownOptions([]byte(text), renderer, options)))
}
