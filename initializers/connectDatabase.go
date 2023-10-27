package initializers

import (
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	var connString string = os.Getenv("DB")
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		panic("Failed Connect to DB")
	}
	DB = db
	// fmt.Println(DB)

}
