package models

import (
	"github.com/faithcomesbyhearing/language-api/language/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var dsn = util.Getenv("MYSQL_CONNECT_STRING", "root:password@tcp(host.docker.internal:3306)/LANGUAGE")
	//log.Printf("MYSQL_CONNECT_STRING: %s", dsn)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Language{})

	DB = database
}
