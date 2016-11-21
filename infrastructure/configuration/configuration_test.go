package configuration

import "testing"

func TestReadConfiguration(t *testing.T) {
	_, err := readConfiguration("./conf/tracker.conf")

	if err != nil {
		t.Errorf("Read configuration failed : %s\n", err)
	}
	if GetDbHost() != "127.0.0.1" {
		t.Errorf("Read host fail got %s expected %s\n", GetDbHost(), "127.0.0.1")
	}

	if GetPort() != 5432 {
		t.Errorf("Read port fail got %d expected %d\n", GetPort(), 5432)
	}

	if GetDatabase() != "tracker_test" {
		t.Errorf("Read database fail got %s expected %s\n", GetDatabase(), "tracker_test")
	}
	if GetDbUser() != "seb" {
		t.Errorf("Read user fail got %s expected %s\n", GetDbUser(), "seb")
	}
	if GetDbPassword() != "passw0rd" {
		t.Errorf("Read password fail got %s expected %s\n", GetDbPassword(), "passw0rd")
	}
}

func TestReadConfigurationWithErrors(t *testing.T) {
	_, err := readConfiguration("inexistant.sample")

	if err == nil {
		t.Error("Read configuration from inexistant.sample should produce error")
	}
}
