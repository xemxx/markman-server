package main

import (
	"log"
	"markman-server/model"
	"markman-server/router"
	"markman-server/tools/config"
	"net/http"

	"golang.org/x/exp/slog"
)

func main() {
	config.Init()
	err := model.Init()
	if err != nil {
		log.Fatal(err)
	}
	r := router.InitRouter()
	cfg := config.Cfg

	s := &http.Server{
		Addr:           ":" + cfg.GetString("server.http_port"),
		Handler:        r,
		ReadTimeout:    cfg.GetDuration("server.read_timeout"),
		WriteTimeout:   cfg.GetDuration("server.write_timeout"),
		MaxHeaderBytes: 1 << 20,
	}

	err = s.ListenAndServe()
	if err != nil {
		slog.Error("启动失败, error: %v", err)
	}
}
