package middleware

import (
	"cross-check/check"
	"cross-check/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	if strings.EqualFold(model.DATABASE, "pgsql") {
		if check.PgsqlStatus.Error != nil {
			http.Error(w, "500 Internal Server Error.\nError: ("+check.PgsqlStatus.Error.Error()+")", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if check.PgsqlStatus.IsMaster {
			_, err := w.Write([]byte("PostgreSQL master is running.\n"))
			if err != nil {
				log.Error(err)
			}
		} else {
			_, err := w.Write([]byte("PostgreSQL slave is running.\n"))
			if err != nil {
				log.Error(err)
			}
		}

		if check.PgsqlStatus.IsWritable {
			_, err := w.Write([]byte("PostgreSQL is writable.\n"))
			if err != nil {
				log.Error(err)
			}
		} else {
			_, err := w.Write([]byte("PostgreSQL is NOT writable.\n"))
			if err != nil {
				log.Error(err)
			}
		}
	} else if strings.EqualFold(model.DATABASE, "mysql") {
		if check.MysqlStatus.Error != nil {
			http.Error(w, "500 Internal Server Error.\nError: ("+check.MysqlStatus.Error.Error()+")", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if check.MysqlStatus.IsMaster {
			_, err := w.Write([]byte("MySQL master is running.\n"))
			if err != nil {
				log.Error(err)
			}
		} else {
			_, err := w.Write([]byte("MySQL slave is running. (Slave lag: " + check.MysqlRows.SecondsBehindMaster + ")\n"))
			if err != nil {
				log.Error(err)
			}
		}
	} else {
		http.Error(w, "500 Internal Server Error.\nError: (Cant set the database.)", http.StatusInternalServerError)
		return
	}
}
