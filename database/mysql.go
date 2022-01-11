package database

import (
	"cross-check/cfg"
	"cross-check/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func MysqlQuery(query string) (model.MysqlRows, error) {

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Values.Mysql.User, cfg.Values.Mysql.Password, cfg.Values.Mysql.Host, cfg.Values.Mysql.Port, cfg.Values.Mysql.Database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return model.MysqlRows{}, err
	}

	defer db.Close()

	rows, err := db.Query(query)

	if err != nil {
		return model.MysqlRows{}, err
	}

	defer rows.Close()

	status, err := ScanMap(rows)

	if err != nil {
		return model.MysqlRows{}, err
	}

	var returnRows = model.MysqlRows{}
	returnRows.SlaveIORunning = status["Slave_IO_Running"].String
	returnRows.SlaveSQLRunning = status["Slave_SQL_Running"].String
	returnRows.SecondsBehindMaster = status["Seconds_Behind_Master"].String

	return returnRows , err
}


func ScanMap(rows *sql.Rows) (map[string]sql.NullString, error) {

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		err = rows.Err()
		if err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	}

	values := make([]interface{}, len(columns))

	for index := range values {
		values[index] = new(sql.NullString)
	}

	err = rows.Scan(values...)

	if err != nil {
		return nil, err
	}

	result := make(map[string]sql.NullString)

	for index, columnName := range columns {
		result[columnName] = *values[index].(*sql.NullString)
	}

	return result, nil
}