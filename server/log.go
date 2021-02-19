package server

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// InitLog ...ログの出力先を指定します。
func InitLog() {
	os.Mkdir("./log/", 0777)
	f, _ := os.Create("./log/" + time.Now().Format("2006-0102-1504") + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(io.MultiWriter(f))
}
