package configuration

import (
	"os"
	"strings"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	conf, err := readConfiguration(gopath + "/src/github.com/scogneau/tracker/conf/sample.conf")

	if err != nil {
		t.Errorf("Read configuration failed : %s\n", err)
	}
	if conf.db.host != "127.0.0.1" {
		t.Errorf("Read host fail got %s expected %s\n", conf.db.host, "127.0.0.1")
	}

	if conf.db.port != 5432 {
		t.Errorf("Read port fail got %d expected %d\n", conf.db.port, 5432)
	}

	if conf.db.db != "tracker_test" {
		t.Errorf("Read database fail got %s expected %s\n", conf.db.db, "tracker_test")
	}
	if conf.db.user != "user_test" {
		t.Errorf("Read user fail got %s expected %s\n", conf.db.user, "user_test")
	}
	if conf.db.password != "passw0rd" {
		t.Errorf("Read password fail got %s expected %s\n", conf.db.password, "passw0rd")
	}
}

func TestReadWebPort(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	InitFromPath(gopath + "/src/github.com/scogneau/tracker/conf/sample.conf")
	port := GetWebPort()
	expectedPort := 8080
	if port != expectedPort {
		t.Errorf("Web port should be %d, got %d\n", expectedPort, port)
	}
}

func TestReadConfigurationWithErrors(t *testing.T) {
	_, err := readConfiguration("inexistant.sample")

	if err == nil {
		t.Error("Read configuration from inexistant.sample should produce error")
	}
}

func TestInitFromPathInexistant(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code must panic if configuration path file doesn't exist")
		}
	}()
	InitFromPath("inexistant")
}
func TestGetConnectionURLFromEnv(t *testing.T) {
	gopath := os.Getenv("GOPATH")

	InitFromPath(gopath + "/src/github.com/scogneau/tracker/conf/sample.conf")

	dbURL := GetDbConnectionURL()
	expectedURLWithPassword := "user=user_test password=passw0rd dbname=tracker_test sslmode=disable"
	if strings.Compare(expectedURLWithPassword, dbURL) != 0 {
		t.Errorf("PG url connection should be %s got %s instead", expectedURLWithPassword, dbURL)
	}

	c.db.password = ""
	dbURL = GetDbConnectionURL()
	expectedURLWithoutPassword := "user=user_test dbname=tracker_test sslmode=disable"
	if strings.Compare(expectedURLWithoutPassword, dbURL) != 0 {
		t.Errorf("PG url connection should be %s got %s instead", expectedURLWithoutPassword, dbURL)
	}

	c.dbEnv = "PG_ENV"
	envExpectedURL := "expectedUrl"
	os.Setenv(c.dbEnv, envExpectedURL)
	defer os.Unsetenv(c.dbEnv)
	dbURL = GetDbConnectionURL()
	if strings.Compare(dbURL, envExpectedURL) != 0 {
		t.Errorf("PG url connection should be %s got %s instead", envExpectedURL, dbURL)
	}
}
