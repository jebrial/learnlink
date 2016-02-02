package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	//"gopkg.in/mgo.v2" // may try mongo out for this
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
