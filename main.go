package main

import (
	"fmt"
	"markman-server/routers"
	"markman-server/tools/config"
)

func main() {
	fmt.Println(config.Cfg.Get("database"))
	r := routers.InitRouter()
	r.Run(":" + config.Cfg.GetString("server.http_port"))

	// s := &http.Server{
	// 	Addr:           ":" + config.Cfg.GetString("server.http_port"),
	// 	Handler:        r,
	// 	ReadTimeout:    config.Cfg.GetDuration("server.read_timeout"),
	// 	WriteTimeout:   config.Cfg.GetDuration("server.write_timeout"),
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
