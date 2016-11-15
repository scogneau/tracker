package db

import (
	"database/sql"
	"fmt"

	//loading postgresql driver
	_ "github.com/lib/pq"
)

type connectionStatus int

//Connection handle database connexion
type Connection struct {
	*sql.DB
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}
}

//NewSQLConnection create a connection to SQL database
func NewSQLConnection(dbuser, dbPassword, dbName string) (Connection, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbuser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	s := Connection{db}
	return s, err
}

func (c Connection) doInTransaction(queryFunction func(tx *sql.Tx, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error) {
	transaction, err := c.Begin()
	defer transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	result, err := queryFunction(transaction, params)
	if err != nil {
		transaction.Rollback()
	}
	return result, err
}

func (c Connection) doWithoutTransaction(queryFunction func(c Connection, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error) {
	return queryFunction(c, params)
}
