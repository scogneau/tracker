package people

import (
	"database/sql"
	"fmt"

	"github.com/scogneau/tracker/infrastructure/db"
)

//People represents a person
type People struct {
	id        int64
	lastName  string
	firstName string
}

func (p People) String() string {
	return fmt.Sprintf("id %d lastname : %s firstname: %s\n", p.id, p.lastName, p.firstName)
}

//ToJSON return json string representing people
func ToJSON(p People) string {
	return fmt.Sprintf("{id:%d,firstname:%s,lastname:%s}", p.id, p.firstName, p.lastName)
}

//ReadPeopleByID read people with the given id from database
func ReadPeopleByID(id int) (People, error) {
	c, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	f := func(tx *sql.Tx, params ...interface{}) (interface{}, error) {

		query := "SELECT * from people WHERE ID = $1;"
		r := tx.QueryRow(query, id)
		p := People{}
		err = r.Scan(&p.id, &p.lastName, &p.firstName)
		if err != nil {
			return nil, err
		}
		return p, nil

	}

	people, err := c.DoInTransaction(f)
	if err != nil {
		return People{}, err
	}
	return people.(People), err
}
