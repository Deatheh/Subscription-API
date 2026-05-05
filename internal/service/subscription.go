package service

import (
	"subscription/internal/entities"
	"subscription/internal/repository/db"

	utils "subscription/pkg"
)

type SubscriptionService struct {
	repository *db.DatabaseRepository
}

func (sub *SubscriptionService) Add(subscription *entities.Subscription) (*entities.Subscription, error) {
	outSubscription, err := sub.repository.AddSubscription(subscription)
	if err != nil {
		return nil, err
	}
	outSubscription.StartDate, err = utils.DateToMonthYear(outSubscription.StartDate)
	if err != nil {
		return nil, err
	}
	if outSubscription.EndDate != "" {
		outSubscription.EndDate, err = utils.DateToMonthYear(outSubscription.EndDate)
		if err != nil {
			return nil, err
		}
	}
	return outSubscription, nil
}
