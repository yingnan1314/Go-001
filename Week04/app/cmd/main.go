package main

import (
	_ "log"
	"ppwords/app/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8010") // listen and serve on 0.0.0.0:8080
}


