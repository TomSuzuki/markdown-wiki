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

// WritePage ...新規作成ページに必要な情報です。
type WritePage struct {
	Word string
}

// PageStatus ...ページの情報を返すときに使います。
type PageStatus struct {
	Exist bool `json:"is_exist,bool"`
}

// SearchPage ...検索ページの表示に必要な情報。
type SearchPage struct {
	Keyword  string
	WordList []string
}
