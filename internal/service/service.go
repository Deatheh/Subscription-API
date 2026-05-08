package service

import (
	"subscription/internal/config"
	"subscription/internal/entities"
	"subscription/internal/repository"
)

type Subscription interface {
	Add(subscription *entities.Subscription) (*entities.Subscription, error)
	GetAll() (*[]entities.Subscription, error)
	GetById(id int) (*entities.Subscription, error)
	GetSumByFilters(filters *entities.SubscriptionsFilters) (int, error)
	Update(subscription *entities.Subscription) (*entities.Subscription, error)
	Delete(id int) error
}

type Service struct {
	Subscription
}

func NewService(repo *repository.Repository, envConf *config.Config) *Service {
	return &Service{
		Subscription: &SubscriptionService{repository: repo.DatabaseRepository},
	}
}
