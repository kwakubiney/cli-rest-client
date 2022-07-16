package services

import (

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
)

func (u *UserService) UpdateUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage.",
		})
		return
	}
	
	whereClause := c.Query("username")
	if whereClause == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage",
		})
		return
	}

	err = u.r.UpdateUserByUsername(whereClause, user)
	if err != nil{
		serr, isUserNotFoundError := err.(*repository.UserNotFoundError)
	if isUserNotFoundError{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": serr.Error(),
		})
		return
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not read user record from database.",
		})
		return
	}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully updated",
		"user":    user})
}
