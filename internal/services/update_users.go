package services

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
)

func (u *UserService) UpdateUser(c *gin.Context) {
	log.Println("=> PUT https://localhost/user\n" +
		"<=")
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		flag.Usage()
		c.Status(http.StatusBadRequest)
		return
	}

	whereClause := u.c.Options.Where
	err = u.r.UpdateUserByUsername(whereClause, user)
	if err != nil {
		log.Println(err)
		flag.Usage()
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully updated",
		"user":    user})
}
