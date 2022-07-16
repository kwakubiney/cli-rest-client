package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserService) FilterUser(c *gin.Context) {
	by := u.c.Options.By
	whereClause := c.Query(by)
	if by == "" || whereClause == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage",
		})
		return
	}

	filteredUser, err := u.r.FilterUser(by, whereClause)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully filtered",
		"user":    filteredUser})
}
