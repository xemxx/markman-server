package main

import (
	"github.com/gin-gonic/gin"
	"io"

	"markman-server/router"
	"markman-server/tools/config"
	"markman-server/tools/logs"
	"net/http"
	"os"
	"time"
)

func main() {
	r := router.InitRouter()
	cfg := config.Cfg
	// 记录到文件。
	f, err := os.Create(cfg.GetString("runtime.log_url") + "gin.log")
	if err != nil {
		logs.Log(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	s := &http.Server{
		Addr:           ":" + cfg.GetString("server.http_port"),
		Handler:        r,
		ReadTimeout:    cfg.GetDuration("server.read_timeout") * time.Second,
		WriteTimeout:   cfg.GetDuration("server.write_timeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = s.ListenAndServe()
	if err != nil {
		logs.Log("启动失败，error：" + err.Error())
	}
}
