package models

import (
	"browser-manage/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	mysqlConf := config.GetMysqlConf()
	engine, err = xorm.NewEngine("mysql", mysqlConf.DB_USERNAME+":"+mysqlConf.DB_PASSWORD+"@tcp("+mysqlConf.DB_HOST+":"+mysqlConf.DB_PORT+")/"+mysqlConf.DB_DATABASE+"?charset="+mysqlConf.CHARSET)
	if err != nil {
		fmt.Println("Connect to mysql error", err)
		return
	}
	engine.SetMaxIdleConns(mysqlConf.SetMaxIdleConns)
	engine.SetMaxOpenConns(mysqlConf.SetMaxOpenConns)
}
