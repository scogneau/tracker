package people

import "testing"

func TestReadPeopleById(t *testing.T) {

	id := 1
	p, err := ReadPeopleByID(id)
	if p.id != 1 {
		t.Errorf("People id must be %d got %d\n", id, p.id)
	}
	if err != nil {
		t.Errorf("Error while reading people with id %d : %s\n", id, err)
	}
}
