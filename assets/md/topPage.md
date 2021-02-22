# はじめに
ようこそ。マークダウン形式でページを編集できるツールです。
サーバーへのアクセスを制限することでローカルWikiとして利用できます。

## 起動方法
`Golang`を使用します。バージョンは`go 1.15`を使用します。  
以下のコマンドで実行してください。
```go
go run main.go
```

# 使い方
新規作成タブから新しいページを作成します。  
作成した記事は、そのページを開いた状態で編集タブを押すことで書き換えることができます。  
書式は基本的なマークダウンの形式です。

# 使用しているマークダウン生成関数
[TomSuzuki/gomarkdown: replaces markdown strings with HTML strings.](https://github.com/TomSuzuki/gomarkdown)


