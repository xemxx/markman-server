package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/exp/slog"

	_ "markman-server/docs"
	"markman-server/model"
	"markman-server/router"
	"markman-server/tools/config"
)

var (
	configFile = flag.String("c", "app.yaml", "config file path")
)

//	@title			Markman API
//	@version		1.0
//	@description	This is a markman server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	xem
//	@contact.url	https://xemxx.cn
//	@contact.email	xemxx@qq.com

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.basic	BasicAuth
func main() {
	flag.Parse()
	config.Init(*configFile)
	err := model.Init()
	if err != nil {
		log.Fatal(err)
	}
	r := router.InitRouter()
	cfg := config.Cfg

	s := &http.Server{
		Addr:           ":" + cfg.Server.HttpPort,
		Handler:        r,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err = s.ListenAndServe()
	if err != nil {
		slog.Error("启动失败, error: %v", err)
	}
}
