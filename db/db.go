package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// sql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Config struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Name     string `ini:"name"`
}

var DB *gorm.DB

func Open() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", Config.User, Config.Password, Config.Host, Config.Port, Config.Name))
	if err != nil {
		log.Fatal(err)
	}
	DB.DropTableIfExists("users", "repositories")
	DB.AutoMigrate(&User{}, &Repository{})
}
