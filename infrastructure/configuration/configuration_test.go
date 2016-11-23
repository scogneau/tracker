package configuration

import (
	"os"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	c, err := readConfiguration(gopath + "/src/github.com/scogneau/tracker/conf/tracker-test.conf")

	if err != nil {
		t.Errorf("Read configuration failed : %s\n", err)
	}
	if c.db.host != "127.0.0.1" {
		t.Errorf("Read host fail got %s expected %s\n", c.db.host, "127.0.0.1")
	}

	if c.db.port != 5432 {
		t.Errorf("Read port fail got %d expected %d\n", c.db.port, 5432)
	}

	if c.db.db != "tracker_test" {
		t.Errorf("Read database fail got %s expected %s\n", c.db.db, "tracker_test")
	}
	if c.db.user != "seb" {
		t.Errorf("Read user fail got %s expected %s\n", c.db.user, "seb")
	}
	if c.db.password != "passw0rd" {
		t.Errorf("Read password fail got %s expected %s\n", c.db.password, "passw0rd")
	}
}

func TestReadConfigurationWithErrors(t *testing.T) {
	_, err := readConfiguration("inexistant.sample")

	if err == nil {
		t.Error("Read configuration from inexistant.sample should produce error")
	}
}
