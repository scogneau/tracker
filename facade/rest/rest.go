package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/scogneau/tracker/domain/people"
)

func HandlePeopleRead(w http.ResponseWriter, r *http.Request) {
	splits := strings.Split(r.URL.Path, "/")
	idPeople, err := strconv.Atoi(splits[len(splits)-1])
	if err != nil {
		log.Print(err)
	}
	p, err := people.ReadPeopleByID(idPeople)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, "%s", people.ToJSON(p))
}
