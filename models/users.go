package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	engine, err := xorm.NewEngine("mysql", "root@localhost")
    if err != nil {
        panic(err)
    }
}


