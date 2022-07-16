package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
)

func (u *UserService) UpdateUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	whereClause := u.c.Options.Where
	err = u.r.UpdateUserByUsername(whereClause, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "user could not be updated. check usage",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully updated",
		"user":    user})
}
