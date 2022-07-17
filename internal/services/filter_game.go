package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
)

type FilterGameRequest struct {
	By       string `json:"by" binding:"required"`
}

func (g *GameService) FilterGame(c *gin.Context) {

	var filterGameRequest FilterGameRequest
	err := c.BindJSON(&filterGameRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage.",
		})
		return
	}

	by := filterGameRequest.By
	whereClause := c.Query(by)
	if by == "" || whereClause == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage",
		})
		return
	}

	filteredGame, err := g.r.FilterGame(by, whereClause)
	if err != nil{
		serr, isGameNotFoundError := err.(*repository.GameNotFoundError)
	if isGameNotFoundError{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": serr.Error(),
		})
		return
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not read game record from database.",
		})
		return
	}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "game successfully filtered",
		"user":    filteredGame})
}
