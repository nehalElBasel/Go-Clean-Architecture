package repository

import "database/sql"

func ConnectMysql() (*sql.DB, error) {
	return sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_crud")
}
