package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"markman-server/router"
	"markman-server/tools/config"
	"net/http"
	"os"
	"time"
)

func main() {
	r := router.InitRouter()
	cfg := config.Cfg
	// 记录到文件。
	f, _ := os.Create(cfg.GetString("runtime.log_url")+"gin.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)


	s := &http.Server{
		Addr:           ":" + cfg.GetString("server.http_port"),
		Handler:        r,
		ReadTimeout:    cfg.GetDuration("server.read_timeout") * time.Second,
		WriteTimeout:   cfg.GetDuration("server.write_timeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Println("启动失败，error：", err)
	}
}
