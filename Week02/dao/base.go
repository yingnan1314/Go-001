package dao

import (
	"database/sql"
	"fmt"
	"log"
	"qyn/pkg/setting"
)

var (
	DB *sql.DB
)

func init() {

	mysql,err := setting.Cfg.GetSection("database")
	if err != nil {
		panic("database connect fail!")
	}

	mysqlConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s",
		mysql.Key("USER").String(),
		mysql.Key("PASSWORD").String(),
		mysql.Key("HOST").String(),
		mysql.Key("PORT").String(),
		mysql.Key("NAME").String(),
		mysql.Key("PARSETIME").String())
	DB, err = sql.Open("mysql", mysqlConfig)
	if err != nil {
		log.Fatalln(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalln(err)
	}
}
