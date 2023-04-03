package infrastructure

import (
	"github.com/jmoiron/sqlx"
)

import _ "github.com/go-sql-driver/mysql"

func CreateLocalDataBase(dataSource string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}

func populateDB(db *sqlx.DB) error {
	return nil
}
