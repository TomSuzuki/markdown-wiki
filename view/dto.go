package view

import "html/template"

// PageData ...表示用。
type PageData struct {
	MenuInfo MenuInfo
	HTML     template.HTML
}

// MenuInfo ...メニューバー用。現在の位置を示すため。
type MenuInfo struct {
	MenuTop     bool
	MenuSearch  bool
	MenuNewPage bool
	MenuEdit    bool
}

// ErrorPage ...エラーページの表示に必要な情報です。
type ErrorPage struct {
	ErrorCode    string
	ErrorMessage string
}

// WordPage ...単語のページを表示するのに必要な情報です。
type WordPage struct {
	Word         string
	MarkdownText string
	MarkdownHTML template.HTML
}
