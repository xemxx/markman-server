package main

import (
	"markman-server/router"
	"markman-server/tools/config"
	"net/http"
	"time"
)

func main() {

	//fmt.Println(config.Cfg.Get("database"))
	r := router.InitRouter()
	//r.Run(":" + config.Cfg.GetString("server.http_port"))
	//model.Test()
	s := &http.Server{
		Addr:           ":" + config.Cfg.GetString("server.http_port"),
		Handler:        r,
		ReadTimeout:    config.Cfg.GetDuration("server.read_timeout") * time.Second,
		WriteTimeout:   config.Cfg.GetDuration("server.write_timeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
