package service

import (
	"github.com/suaas21/pathao/logger"
	"github.com/suaas21/pathao/repo"
)

type User struct {
	userRepo repo.UserRepository
	log      logger.StructLogger
}

func NewUser(userRepo repo.UserRepository, lgr logger.StructLogger) *User {
	return &User{
		userRepo: userRepo,
		log:      lgr,
	}
}
