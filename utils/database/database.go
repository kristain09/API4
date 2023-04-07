package database

import (
	"fmt"
	"log"

	"github.com/kristain09/API4/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c config.AppConfig) *gorm.DB {
	DBConfig := config.IniConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		DBConfig.DBUSERNAME,
		DBConfig.DBPASSWORD,
		DBConfig.DBHOST,
		DBConfig.DBPORT,
		DBConfig.DBNAME,
		DBConfig.DBARGS,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot connect to database")
	}
	return db
}
