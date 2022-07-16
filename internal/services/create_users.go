package services

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Age      string `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (u *UserService) CreateUser(c *gin.Context) {
	var newUser model.User
	var createUserRequest CreateUserRequest
	err := c.BindJSON(&createUserRequest)
	if err != nil {
		log.Println(err)
		flag.Usage()
		c.Status(http.StatusBadRequest)
		return
	}

	newUser.Age = createUserRequest.Age
	newUser.Email = createUserRequest.Email
	newUser.Username = createUserRequest.Username

	err = u.r.CreateUser(newUser)
	if err != nil {
		log.Println(err)
		flag.Usage()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "user could not be created. check usage.",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "user successfully created",
		"user":    newUser})
}
