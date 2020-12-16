// Package conf 提供最基础的配置加载功能
package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg *ini.File
	err error
)

func init () {
	Cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatal(err)
	}
}