package configuration

import (
	"bufio"
	"fmt"
	"log"
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
	dbEnvKey      = "db.env"
	webPortKey    = "web.port"
)

//Conf contains configuration
var c configuration

//InitFromPath initialize configuration from path
func InitFromPath(path string) {
	var err error
	c, err = readConfiguration(path)
	log.Info(fmt.Sprintf("Reading configuration from %s\n", path))
	if err != nil {
		panic(fmt.Sprintf("Error while reading configuration file :%s\n", err))
	}
}

//Configuration contains configuration for application
type configuration struct {
	db         dbConfiguration
	dbEnv      string
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

//GetWebPort return the port for web application
func GetWebPort() int {
	return c.serverPort
}

//GetDbConnectionURL return  a url to connect to postgresql from parameter or environment
func GetDbConnectionURL() string {
	var dbinfo string

	if c.dbEnv != "" {
		if containsDbConfiguration() {
			log.Print("INFO - Configuration file contains environment setup , all others setup for db will be ignored (host,user,password,dbname,port)")
		}
		return os.Getenv(c.dbEnv)
	} else if c.db.password == "" {
		dbinfo = fmt.Sprintf("user=%s dbname=%s sslmode=disable",
			c.db.user, c.db.db)
	} else {
		dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			c.db.user, c.db.password, c.db.db)
	}
	return dbinfo
}

func containsDbConfiguration() bool {
	return c.db.db != "" || c.db.host != "" || c.db.password != "" || c.db.port != 0 || c.db.user != ""
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
	fmt.Println(rawConfiguration)
	fmt.Println(os.Getenv(rawConfiguration[dbEnvKey]))
	return configuration{
		dbConfiguration{
			host:     rawConfiguration[dbHostKey],
			port:     dbport,
			db:       rawConfiguration[dbNameKey],
			user:     rawConfiguration[dbUserKey],
			password: rawConfiguration[dbPasswordKey],
		},
		rawConfiguration[dbEnvKey],
		webport,
	}, err
}
