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
	dbHostKey     = "db.host"
	dbPortKey     = "db.port"
	dbNameKey     = "db.name"
	dbUserKey     = "db.user"
	dbPasswordKey = "db.password"
	webPortKey    = "web.port"
)

//Conf contains configuration
var c configuration

func init() {
	path := flag.String("c", "conf/tracker.conf", "Path of configuration file")
	var err error
	c, err = readConfiguration(*path)
	if err != nil {
		panic(fmt.Sprintf("Error while reading configuration file :%s\n", err))
	}
}

//Configuration contains configuration for application
type configuration struct {
	db         dbConfiguration
	serverPort int
}

//DbConfiguration contains db configuration
type dbConfiguration struct {
	host     string
	port     int
	db       string
	user     string
	password string
}

//GetDbHost return database host
func GetDbHost() string {
	return c.db.host
}

//GetPort return dabase port
func GetPort() int {
	return c.db.port
}

//GetDatabase return database name
func GetDatabase() string {
	return c.db.db
}

//GetDbUser return user use to connect to database
func GetDbUser() string {
	return c.db.user
}

//GetDbPassword return password use to connect to database
func GetDbPassword() string {
	return c.db.password
}

//readConfiguration read configuration from file
func readConfiguration(path string) (configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return configuration{}, err
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

	return configuration{
		dbConfiguration{
			host:     rawConfiguration[dbHostKey],
			port:     dbport,
			db:       rawConfiguration[dbNameKey],
			user:     rawConfiguration[dbUserKey],
			password: rawConfiguration[dbPasswordKey],
		},
		webport,
	}, err
}
