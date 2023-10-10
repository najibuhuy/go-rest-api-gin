package model

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	const connString string = "sqlserver://backend_service:A!s3d4f5g6h7j8k9l0@localhost:1433?database=go-gin-restapi"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	DB = db
}
