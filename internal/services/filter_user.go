package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
)

type FilterUserRequest struct {
	By       string `json:"by" binding:"required"`
}

func (u *UserService) FilterUser(c *gin.Context) {

	var filterUserRequest FilterUserRequest
	err := c.BindJSON(&filterUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage.",
		})
		return
	}

	by := filterUserRequest.By
	whereClause := c.Query(by)
	if by == "" || whereClause == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage",
		})
		return
	}

	filteredUser, err := u.r.FilterUser(by, whereClause)
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
		"message": "user successfully filtered",
		"user":    filteredUser})
}
