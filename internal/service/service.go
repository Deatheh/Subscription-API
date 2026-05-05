package service

import (
	"subscription/internal/config"
	"subscription/internal/repository"
)

type Service struct {
}

func NewService(repository *repository.Repository, envConf *config.Config) *Service {
	return &Service{}
}
