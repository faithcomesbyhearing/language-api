package models

import (
	"log"
	"os"
	"time"

	"github.com/faithcomesbyhearing/language-api/language/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	var dsn = util.Getenv("MYSQL_CONNECT_STRING", "root:password@tcp(host.docker.internal:3306)/LANGUAGE?charset=utf8mb4&parseTime=True&loc=Local")
	//log.Printf("MYSQL_CONNECT_STRING: %s", dsn)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
