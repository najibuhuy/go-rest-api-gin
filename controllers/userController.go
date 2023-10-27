package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/najibuhuy/go-rest-api-gin/dto"
	"github.com/najibuhuy/go-rest-api-gin/initializers"
	"github.com/najibuhuy/go-rest-api-gin/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//get the email and pass of req.body

	if c.Bind(&dto.SignUpUserRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body req"})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.SignUpUserRequest.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	//create user
	user := models.User{Email: dto.SignUpUserRequest.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create User"})
		return
	}
	//response

	responses := dto.SignUpUserResponse{
		Status:  http.StatusOK,
		Data:    user,
		Message: "SUCCES_GET_DATA",
	}
	c.IndentedJSON(http.StatusOK, responses)
	// c.JSON(http.StatusOK, gin.H{responses})
}

func Login(c *gin.Context) {
	// get the email and pass from req.body
	if c.Bind(&dto.SignUpUserRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body req"})
		return
	}

	//look up requsted user
	var user models.User
	fmt.Println(dto.LoginUserRequest.Email, "Email")
	initializers.DB.First(&user, "email = ?", dto.LoginUserRequest.Email)

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password"})
		return
	}

	//compare sent in pass with saved pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.LoginUserRequest.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decrypt password"})
		return
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	fmt.Println(err, "err")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func MidTest(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"data": user})

}
