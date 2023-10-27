package dto

import "github.com/najibuhuy/go-rest-api-gin/models"

type ListProductResponse struct {
	Status  int              `json:"status"`
	Data    []models.Product `json:"data"`
	Message string           `json:"message"`
}
