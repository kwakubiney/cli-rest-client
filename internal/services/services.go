package services

import (
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
	"github.com/kwakubiney/canonical-take-home/internal/handlers/cli"
)

type UserService struct {
	r repository.UserRepository
	c cli.CliHandler
}

func NewUserService(repository *repository.UserRepository, cliHandler *cli.CliHandler) *UserService {
	return &UserService{
		r: *repository,
		c: *cliHandler,
	}
}
