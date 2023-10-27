package dto

import "github.com/najibuhuy/go-rest-api-gin/models"

var SignUpUserRequest struct {
	Email    string
	Password string
}

var LoginUserRequest struct {
	Email    string
	Password string
}

type SignUpUserResponse struct {
	Status  int         `json:"status"`
	Data    models.User `json:"data"`
	Message string      `json:"message"`
}
