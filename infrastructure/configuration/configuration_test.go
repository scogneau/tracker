package configuration

import "testing"

func TestReadConfiguration(t *testing.T) {
	c, err := ReadConfiguration("./conf/tracker.conf")

	if err != nil {
		t.Errorf("Read configuration failed : %s\n", err)
	}
	if c.Db.GetHost() != "127.0.0.1" {
		t.Errorf("Read host fail got %s expected %s\n", c.Db.GetHost(), "127.0.0.1")
	}

	if c.Db.GetPort() != 5432 {
		t.Errorf("Read host fail got %d expected %d\n", c.Db.GetPort(), 5432)
	}

	if c.Db.getDatabase() != "tracker_test" {
		t.Errorf("Read host fail got %s expected %s\n", c.Db.getDatabase(), "tracker_test")
	}
}

func TestReadConfigurationWithErrors(t *testing.T) {
	_, err := ReadConfiguration("inexistant.sample")

	if err == nil {
		t.Error("Read configuration from inexistant.sample should produce error")
	}
}
