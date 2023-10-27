package initializers

import "github.com/najibuhuy/go-rest-api-gin/models"

func AutoMigrateModel() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})

}
