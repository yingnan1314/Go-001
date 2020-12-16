package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"ppwords/util/conf"
	"strings"
)

var (
	DB *sql.DB
	err error
)

func init() {
	mysql := conf.Cfg.Section("mysql").KeysHash()

	mysqlConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s",
		mysql["username"],
		mysql["password"],
		mysql["host"],
		mysql["port"],
		mysql["database"],
		mysql["parseTime"])
	DB, err = sql.Open("mysql", mysqlConfig)
	if err != nil {
		log.Fatalln(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalln(err)
	}
}

func Placeholders(cnt int) string {
	var str strings.Builder
	if cnt <= 0 {
		return ""
	}
	for i :=0; i < cnt - 1; i++ {
		str.WriteString("?,")
	}
	str.WriteString("?")
	return str.String()
}