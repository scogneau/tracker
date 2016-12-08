package db

import (
	"database/sql"
	"fmt"
	"log"

	//loading postgresql driver
	_ "github.com/lib/pq"
	"github.com/scogneau/tracker/infrastructure/configuration"
)

//Connection handle database connexion
type Connection struct {
	*sql.DB
}

func checkErr(err error) {
	if err != nil {
		log.Print("Error", err)
		panic(err)
	}
}

//NewSQLConnection create a connection to SQL database
func newSQLConnection(dbuser, dbPassword, dbName string) (Connection, error) {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		dbuser, dbName)
	log.Print(dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	log.Print(err)
	s := Connection{db}
	return s, err
}

//Connect create a database connection using configuration information
func Connect() (Connection, error) {
	return newSQLConnection(configuration.GetDbUser(), configuration.GetDbPassword(), configuration.GetDatabase())
}

//DoInTransaction execute queryFunction in a transaction
func (c Connection) DoInTransaction(queryFunction func(tx *sql.Tx, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error) {
	transaction, err := c.Begin()
	if err != nil {
		fmt.Println(err, transaction)
		return 0, err
	}

	defer transaction.Commit()

	result, err := queryFunction(transaction, params)
	if err != nil {
		transaction.Rollback()
	}
	return result, err
}

//DoWithoutTransaction execute queryFunction without transaction
func (c Connection) DoWithoutTransaction(queryFunction func(c Connection, parameters ...interface{}) (interface{}, error), params ...interface{}) (interface{}, error) {
	return queryFunction(c, params)
}
