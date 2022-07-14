package services

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
)

func (u *UserService) FindUser(c *gin.Context){
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
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "user successfully created",
		"user": newUser})
}
