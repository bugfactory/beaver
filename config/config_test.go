package configuration

import "testing"

func TestNewValidFile(t *testing.T) {
	//New (filename string) (Configuration, error)
	conf, err := New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}

	if conf.Ip != "123.123.123.21" {
		t.Fatal("error")
	}
}

func TestNewInValidFile(t *testing.T) {
	_, err := New("./xxxxx.yml")
	if err == nil {
		t.Fatal(err)
	}
}


