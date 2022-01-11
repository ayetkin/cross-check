package check

import (
	"cross-check/database"
	log "github.com/sirupsen/logrus"
)

func PgsqlCheck(){
	isSlave, err := database.PgsqlQuery("SELECT pg_is_in_recovery()")
	if err != nil {
		PgsqlStatus.Error=err
		log.Error(err)
	}else{
		PgsqlStatus.Error=nil
		if isSlave == "false" {
			log.Warning("This node is Master!")
			PgsqlStatus.IsMaster = true
		}else {
			log.Warning("This node is Slave!")
			PgsqlStatus.IsMaster = false
		}
	}

	isReadonly, err := database.PgsqlQuery("SHOW transaction_read_only")
	if err != nil {
		PgsqlStatus.Error=err
		log.Error(err)
	}else{
		PgsqlStatus.Error=nil
		if isReadonly == "off" {
			log.Warning("This node is writable!")
			PgsqlStatus.IsWritable = true
		}else{
			log.Warning("This node is Readonly, NOT writable!")
			PgsqlStatus.IsWritable = false
		}
	}
}