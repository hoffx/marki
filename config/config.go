package config

import (
	"log"

	"github.com/hoffx/marki/db"
	ini "gopkg.in/ini.v1"
)

var Config struct {
}

func Load(path string) {
	f, err := ini.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Section("db").MapTo(&db.Config)
	if err != nil {
		log.Fatal(err)
	}
	err = f.MapTo(&Config)
	if err != nil {
		log.Fatal(err)
	}
}
