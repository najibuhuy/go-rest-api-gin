package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/dto"
	"github.com/najibuhuy/go-rest-api-gin/initializers"
	"github.com/najibuhuy/go-rest-api-gin/models"
)

func GetListProduct(c *gin.Context) {
	channel := make(chan []models.Product) //async
	defer close(channel)                   //after all of job has done

	var data []models.Product
	initializers.DB.Find(&data)

	go func() { //goroutine  skip before this func is done
		// time.Sleep(2 * time.Second)
		channel <- data //await
		fmt.Println("selesai mengirim data ke chnnel")

	}()
	listData := <-channel
	response := dto.ListProductResponse{
		Status:  http.StatusOK,
		Data:    listData,
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
