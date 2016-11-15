package db

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestReadFromDb(t *testing.T) {
	c, err := NewSQLConnection("seb", "seb", "tracker_test")
	if err != nil {
		t.Error(err)
	}

	var selectFnC = func(c Connection, input ...interface{}) (interface{}, error) {

		var resultInt int
		r := c.QueryRow("SELECT 1;")
		r.Scan(&resultInt)
		return resultInt, nil
	}

	r, err := c.doWithoutTransaction(selectFnC, 1, 2)
	if err != nil {
		fmt.Println("Error ", err)
	}
	if r != 1 {
		intValue, _ := r.(int)
		t.Errorf("SELECT 1 should return 1, got %d\n", intValue)
	}

}

func TestReadFromDbWithParameters(t *testing.T) {
	c, err := NewSQLConnection("seb", "seb", "tracker_test")
	if err != nil {
		t.Error(err)
	}

	var selectFnC = func(c Connection, input ...interface{}) (interface{}, error) {

		var resultInt int
		num := 1
		r := c.QueryRow("SELECT count(*) from people where id=$1;", int(num))
		r.Scan(&resultInt)
		return resultInt, nil
	}

	r, err := c.doWithoutTransaction(selectFnC, 1, 2)
	if err != nil {
		fmt.Println("Error ", err)
	}
	if r != 1 {
		intValue, _ := r.(int)
		t.Errorf("SELECT 1 should return 1, got %d\n", intValue)
	}

}

func TestReadWithTransaction(t *testing.T) {
	c, err := NewSQLConnection("seb", "seb", "tracker_test")
	if err != nil {
		t.Error(err)
	}

	var selectFnC = func(tx *sql.Tx, input ...interface{}) (interface{}, error) {

		var resultInt int
		num := 1
		r := tx.QueryRow("SELECT count(*) from people where id=$1;", int(num))
		r.Scan(&resultInt)
		return resultInt, nil

	}

	r, err := c.doInTransaction(selectFnC, 1, 2)
	if err != nil {
		fmt.Println("Error ", err)
	}
	if r != 1 {
		intValue, _ := r.(int)
		t.Errorf("SELECT 1 should return 1, got %d\n", intValue)
	}
}
