package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Execute(dbConnection string, query string, params ...interface{}) sql.Result {
	db, err := sql.Open("postgres", dbConnection)
	checkError(err)
	defer db.Close()
	res, e := db.Exec(query, params...)
	checkError(e)
	return res
}

func Query(dbConnection string, query string, params ...interface{}) *sql.Rows {
	db, err := sql.Open("postgres", dbConnection)
	checkError(err)
	defer db.Close()
	res, e := db.Query(query, params...)
	checkError(e)
	return res
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
