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

func (sub *SubscriptionService) GetAll() (*[]entities.Subscription, error) {
	subscriptionSlice, err := sub.repository.GetAllSubscription()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(*subscriptionSlice); i++ {
		(*subscriptionSlice)[i].StartDate, err = utils.DateToMonthYear((*subscriptionSlice)[i].StartDate)
		if err != nil {
			return nil, err
		}
		if (*subscriptionSlice)[i].EndDate != "" {
			(*subscriptionSlice)[i].EndDate, err = utils.DateToMonthYear((*subscriptionSlice)[i].EndDate)
			if err != nil {
				return nil, err
			}
		}
	}
	return subscriptionSlice, nil
}

func (sub *SubscriptionService) GetById(id int) (*entities.Subscription, error) {
	outSubscription, err := sub.repository.GetSubscriptionById(id)
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

func (sub *SubscriptionService) GetSumByFilters(filters *entities.SubscriptionsFilters) (int, error) {
	sum, err := sub.repository.GetSumSubscriptionByFilters(filters)
	if err != nil {
		return 90, err
	}
	return sum, nil
}

func (sub *SubscriptionService) Update(subscription *entities.Subscription) (*entities.Subscription, error) {
	outSubscription, err := sub.repository.UpdateSubscription(subscription)
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

func (sub *SubscriptionService) Delete(id int) error {
	err := sub.repository.DeleteSubscription(id)
	if err != nil {
		return err
	}
	return nil
}
