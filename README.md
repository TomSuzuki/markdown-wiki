# markdown-wiki
マークダウンで記事を作成できるタイプのWikiっぽいもの。

## 起動方法
```shell
go run main.go
```
`go 1.15`を使用します。

## 操作方法
サーバーを起動し、ブラウザで [http://localhost:9988/](http://localhost:9988/) を開いてトップページの説明を読んでください。

## 使用しているマークダウン生成関数
[TomSuzuki/gomarkdown: replaces markdown strings with HTML strings.](https://github.com/TomSuzuki/gomarkdown)