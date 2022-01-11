package check

import (
	"cross-check/database"
	log "github.com/sirupsen/logrus"
)

func MysqlCheck() {
	rows, err := database.MysqlQuery("SHOW SLAVE STATUS")
	if err != nil {
		MysqlStatus.Error = err
		log.Error(err)
	} else {
		MysqlStatus.Error = nil
		if rows.SlaveSQLRunning == "Yes" && rows.SlaveIORunning == "Yes" {
			log.Warning("MySQL slave is running. (Slave lag: ", rows.SecondsBehindMaster, ")")
			MysqlStatus.IsMaster = false
		} else {
			log.Warning("MySQL master is running.")
			MysqlStatus.IsMaster = true
		}
	}
}
