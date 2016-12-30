package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/scogneau/tracker/facade/rest"
	"github.com/scogneau/tracker/infrastructure/configuration"
)

func main() {
	path := flag.String("c", "conf/tracker.conf", "Path of configuration file")
	flag.Parse()
	configuration.InitFromPath(*path)
	http.HandleFunc("/people/", rest.HandlePeopleRead)
	fmt.Println("Starting tracker application")

	port := os.Getenv("PORT")
	if port == "" {
		log.Print("No $PORT defined fallback to configuration")
		port = strconv.Itoa(configuration.GetWebPort())
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
