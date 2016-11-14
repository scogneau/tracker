package db

import (
	"errors"
	"fmt"
	"testing"
)

func TestReadFromDb(t *testing.T) {
	c, err := NewSQLConnection("seb", "seb", "tracker_test")
	if err != nil {
		t.Error(err)
	}

	var selectFnC = func(c Connection, input ...interface{}) (interface{}, error) {

		if sqlC, ok := c.(sqlConnection); ok {
			var resultInt int
			r := sqlC.QueryRow("SELECT 1;")
			r.Scan(&resultInt)
			return resultInt, nil
		}

		return nil, errors.New("La connection doit être une connection SQL")
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

		if sqlC, ok := c.(sqlConnection); ok {
			var resultInt int
			num := 1
			r := sqlC.QueryRow("SELECT count(*) from people where id=$1;", int(num))
			r.Scan(&resultInt)
			return resultInt, nil
		}

		return nil, errors.New("La connection doit être une connection SQL")
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
