package services

import (
	//"log"

	"github.com/gin-gonic/gin"
	//"github.com/kwakubiney/canonical-take-home/internal/domain/model"
)

// find all
func (u *UserService) FindUsers(c *gin.Context){
	c.JSON(200, gin.H{"message": u.c.Options.Fields})
}




