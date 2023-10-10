package productscontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/models"
)

type Response struct {
	Status  int              `json:"status"`
	Data    []models.Product `json:"data"`
	Message string           `json:"message"`
}

func GetListProduct(c *gin.Context) {
	var data []models.Product

	models.DB.Find(&data)
	response := Response{
		Status:  http.StatusOK,
		Data:    data,
		Message: "SUCCES_GET_DATA",
	}
	fmt.Println(response, "response")

	c.IndentedJSON(http.StatusOK, response)
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data, "message": "SUCCES_GET_DATA"})
}
func GetDetailProduct(c *gin.Context) {

}
func CreateProduct(c *gin.Context) {

}
func UpdateProduct(c *gin.Context) {

}
func DeleteProduct(c *gin.Context) {

}
