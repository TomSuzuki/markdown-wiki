# goバージョン
FROM golang:1.19.1-alpine

# アップデートとgitのインストール
RUN apk update && apk add git

# markdown-wikiディレクトリの作成
RUN mkdir -p /go/src/github.com/markdown-wiki

# ワーキングディレクトリの設定
WORKDIR /go/src/github.com/markdown-wiki

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY  . /go/src/github.com/markdown-wiki

# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports

# ビルド
RUN go build -o /app

# 起動
CMD ["/app"]
