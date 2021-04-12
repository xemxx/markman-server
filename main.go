package main

import (
	"markman-server/router"
	"markman-server/tools/config"
	"markman-server/tools/logs"
	"net/http"
	"time"
)

func main() {
	r := router.InitRouter()
	cfg := config.Cfg

	s := &http.Server{
		Addr:           ":" + cfg.GetString("server.http_port"),
		Handler:        r,
		ReadTimeout:    cfg.GetDuration("server.read_timeout") * time.Second,
		WriteTimeout:   cfg.GetDuration("server.write_timeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		logs.Log("启动失败，error：" + err.Error())
	}
}
