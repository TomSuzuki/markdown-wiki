package main

import "github.com/TomSuzuki/markdown-wiki/server"

func main() {
	server.InitLog()
	server.Router().Run(":9988")
}
