package configuration

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	dbHostKey  = "db.host"
	dbPortKey  = "db.port"
	dbNameKey  = "db.name"
	webPortKey = "web.port"
)

//Conf contains configuration
var Conf Configuration

func init() {
	path := flag.String("c", "conf/tracker.conf", "Path of configuration file")
	var err error
	Conf, err = ReadConfiguration(*path)
	if err != nil {
		panic(fmt.Sprintf("Error while reading configuration file :%s\n", err))
	}
}

//Configuration contains configuration for application
type Configuration struct {
	Db         DbConfiguration
	ServerPort int
}

//DbConfiguration contains db configuration
type DbConfiguration struct {
	host string
	port int
	db   string
}

//GetHost return database host
func (db DbConfiguration) GetHost() string {
	return db.host
}

//GetPort return dabase port
func (db DbConfiguration) GetPort() int {
	return db.port
}

//getDatabase return database name
func (db DbConfiguration) getDatabase() string {
	return db.db
}

//ReadConfiguration read configuration from file
func ReadConfiguration(path string) (Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return Configuration{}, err
	}

	rawConfiguration := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if !strings.HasPrefix(t, "#") {
			l := strings.Split(t, ":")
			rawConfiguration[strings.TrimSpace(l[0])] = strings.TrimSpace(l[1])
		}

	}
	webport, err := strconv.Atoi(rawConfiguration[webPortKey])
	dbport, err := strconv.Atoi(rawConfiguration[dbPortKey])
	return Configuration{
		DbConfiguration{
			host: rawConfiguration[dbHostKey],
			port: dbport,
			db:   rawConfiguration[dbNameKey],
		},
		webport,
	}, err
}
