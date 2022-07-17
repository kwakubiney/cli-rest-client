package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/canonical-take-home/internal/services"
)

type Server struct {
	userService *services.UserService
	gameService *services.GameService
	e       *gin.Engine
	srv     http.Server
}

func New(userService *services.UserService, gameService *services.GameService) *Server {
	gin.SetMode(gin.ReleaseMode)
	return &Server{
		userService: userService,
		gameService: gameService,
		e:       gin.Default(),
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	s.e.POST("/User", s.userService.CreateUser)
	s.e.PUT("/User", s.userService.UpdateUser)
	s.e.GET("/User", s.userService.FilterUser)
	s.e.POST("/Game", s.gameService.CreateGame)
	s.e.GET("/Game", s.gameService.FilterGame)
	s.e.PUT("/Game", s.gameService.UpdateGame)
	
	return s.e
}

func (s *Server) Start() {
	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: s.SetupRoutes(),
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := s.srv.Close(); err != nil {
			log.Println("failed to shutdown server", err)
		}
	}()

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("server closed after interruption")
			} else {
				log.Println("unexpected server shutdown. err:", err)
			}
		}
	}()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
