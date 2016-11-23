package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/scogneau/tracker/facade/rest"
	"github.com/scogneau/tracker/infrastructure/configuration"
)

func main() {
	path := flag.String("c", "conf/tracker.conf", "Path of configuration file")
	configuration.InitFromPath(*path)
	fmt.Println("Test")
	http.HandleFunc("/people/", rest.HandlePeopleRead)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.GetWebPort()), nil))
}
