package services

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
)

type CreateGameRequest struct {
	AgeRating  string `json:"age_rating" binding:"required"`
	Title	string   `json:"title" binding:"required"`
	Description	string   `json:"description" binding:"required"`
	Publisher   string   `json:"publisher" binding:"required"`
	URL        string   `json:"url" binding:"required"`
}

func (g *GameService) CreateGame(c *gin.Context) {
	var newGame model.Game
	var createGameRequest CreateGameRequest
	err := c.BindJSON(&createGameRequest)
	if err != nil {
		log.Println(err)
		flag.Usage()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check usage.",
		})
		return
	}

	newGame.AgeRating = createGameRequest.AgeRating
	newGame.URL = createGameRequest.URL
	newGame.Publisher = createGameRequest.Publisher
	newGame.Title = createGameRequest.Title
	newGame.Description = createGameRequest.Description

	err = g.r.CreateGame(newGame)
	if err != nil {
		log.Println(err)
		flag.Usage()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "game could not be created. check usage.",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "game successfully created",
		"game":    newGame})
}
