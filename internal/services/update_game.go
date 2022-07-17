package services

import (

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
)

func (g *GameService) UpdateGame(c *gin.Context) {
	var game model.Game
	err := c.BindJSON(&game)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage.",
		})
		return
	}
	
	whereClause := c.Query("title")
	if whereClause == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage",
		})
		return
	}
	err = g.r.UpdateGameByTitle(whereClause, game)
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
		"message": "game successfully updated",
		"game":    game,
	})
}
