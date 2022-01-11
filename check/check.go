package check

import (
	"cross-check/model"
	"strings"
	"time"
)

var PgsqlStatus = model.Status{}
var MysqlStatus = model.Status{}
var MysqlRows = model.MysqlRows{}

func Retry(database string){
	if strings.EqualFold(database, "pgsql") {
		for {
			<-time.After(1 * time.Second)
			go PgsqlCheck()
		}
	}else if strings.EqualFold(database, "mysql") {
		for {
			<-time.After(1 * time.Second)
			go MysqlCheck()
		}
	}
}