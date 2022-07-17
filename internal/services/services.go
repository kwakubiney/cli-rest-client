package services

import (
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
)

type UserService struct {
	r repository.UserRepository
	c cli.CliHandler
}

type GameService struct {
	r repository.GameRepository
	c cli.CliHandler
}

func NewUserService(userRepository *repository.UserRepository, cliHandler *cli.CliHandler) *UserService {
	return &UserService{
		r: *userRepository,
		c: *cliHandler,
	}
}

func NewGameService(repository *repository.GameRepository, cliHandler *cli.CliHandler) *GameService {
	return &GameService{
		r: *repository,
		c: *cliHandler,
	}
}

