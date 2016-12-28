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
	flag.Parse()
	configuration.InitFromPath(*path)
	http.HandleFunc("/people/", rest.HandlePeopleRead)
	fmt.Println("Starting tracker application")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.GetWebPort()), nil))
}
