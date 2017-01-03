package main

import (
	"flag"
	"fmt"
	"io"
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

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var path string
		if request.RequestURI == "/" {
			path = "static/index.html"
		} else {
			path = "static/" + request.RequestURI
		}

		f, err := os.Open(path)
		if err != nil {
			log.Printf("Error reading index.html %s", err)
		}
		io.Copy(writer, f)
	})
	http.HandleFunc("/people/", rest.HandlePeopleRead)
	fmt.Println("Starting tracker application")

	port := os.Getenv("PORT")
	if port == "" {
		log.Print("No $PORT defined fallback to configuration")
		port = strconv.Itoa(configuration.GetWebPort())
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
