package services

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (u *UserService) FilterUser(c *gin.Context) {
	fmt.Println("Getting user...........")
	fmt.Println("=> GET https://localhost/user\n" +
		"<=")
	by := u.c.Options.By
	whereClause := c.Query(by)
	if by == "" || whereClause == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	filteredUser, err := u.r.FilterUser(by, whereClause)
	if err != nil {
		log.Println(err)
		flag.Usage()
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully filtered",
		"user":    filteredUser})
}
