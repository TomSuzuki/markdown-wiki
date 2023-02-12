<p align="center">
 <img width="100px" src="./assets/favicon.ico" align="center" alt="Icon" />
 <h2 align="center"><a href="https://tomsuzuki.github.io/markdown-wiki/">markdown-wiki</a></h2>
 <p align="center">マークダウンで記事を作成できるタイプのWikiっぽいもの。</p>
</p>

## ビルド＆起動
Dockerを使用します。

### ビルド
```sh
docker-compose build
```

### 起動
```sh
docker-compose up -d
```

## 特徴
- マークダウンファイルとして保存
- Docker上で動作

## 操作方法
サーバーを起動し、ブラウザで [http://localhost:9988/](http://localhost:9988/) を開いてトップページの説明を読んでください。

## 画面
<img src="./readme/sample.png" width="320px">
<img src="./readme/ca7525c4c43782f4161de45169b9f6ce.gif" width="320px">


## 使用しているマークダウン生成関数
[TomSuzuki/white600: Markdown → HTML](https://github.com/TomSuzuki/white600)
