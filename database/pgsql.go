package database

import (
	"cross-check/cfg"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func PgsqlQuery(query string) (string, error) {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Values.Pgsql.Host, cfg.Values.Pgsql.Port, cfg.Values.Pgsql.User, cfg.Values.Pgsql.Password, cfg.Values.Pgsql.Database)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return "", err
	}

	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			return "", err
		}
		return value, nil
	}
	return "Connected to database.", nil
}