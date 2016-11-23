package people

import (
	"os"
	"testing"

	"github.com/scogneau/tracker/infrastructure/configuration"
)

func TestReadPeopleById(t *testing.T) {

	gopath := os.Getenv("GOPATH")
	configuration.InitFromPath(gopath + "/src/github.com/scogneau/tracker/conf/tracker.conf")

	id := 1
	p, err := ReadPeopleByID(id)
	if p.id != 1 {
		t.Errorf("People id must be %d got %d\n", id, p.id)
	}
	if err != nil {
		t.Errorf("Error while reading people with id %d : %s\n", id, err)
	}
}
