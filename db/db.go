package db

import (
	"database/sql"
	"fmt"

	//loading postgresql driver
	_ "github.com/lib/pq"
)

type connectionStatus int

//Connection handle database connexion
type sqlConnection struct {
	*sql.DB
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}
}

//Connection handle connection to database
type Connection interface {
	doInTransaction(queryFunction func(tx *sql.Tx, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error)

	doWithoutTransaction(queryFunction func(c Connection, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error)
}

//NewSQLConnection create a connection to SQL database
func NewSQLConnection(dbuser, dbPassword, dbName string) (Connection, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbuser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	s := &sqlConnection{db}
	return s, err
}

func (s sqlConnection) doInTransaction(queryFunction func(tx *sql.Tx, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error) {
	transaction, err := s.Begin()
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

func (s sqlConnection) doWithoutTransaction(queryFunction func(c Connection, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error) {
	err := s.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return queryFunction(s, params)
}
