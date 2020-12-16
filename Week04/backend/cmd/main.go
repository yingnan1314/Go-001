package main

import (
	_ "log"
	"ppwords/backend/router"
	"ppwords/util/conf"
)

func main() {

	cfg := conf.Cfg.Section("app")
	host := cfg.Key("backend.port").String()

	r := router.InitRouter()
	r.Run(host) // listen and serve on 0.0.0.0:8080
}
