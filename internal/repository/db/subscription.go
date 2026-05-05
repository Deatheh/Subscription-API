package db

import (
	"database/sql"
	"fmt"
	"subscription/internal/entities"
)

func (dbr *DatabaseRepository) AddSubscription(sub *entities.Subscription) (*entities.Subscription, error) {
	var endDate interface{}
	if sub.EndDate == "" {
		endDate = nil
	} else {
		endDate = sub.EndDate
	}

	var scannedEndDate sql.NullString
	err := dbr.DB.QueryRow(`
		INSERT INTO subscription (name, price, user_uuid, date_start, date_end)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, price, user_uuid, date_start, date_end`,
		sub.ServiceName, sub.Price, sub.UserUUID, sub.StartDate, endDate).Scan(
		&sub.Id, &sub.ServiceName, &sub.Price, &sub.UserUUID, &sub.StartDate, &scannedEndDate)
	if err != nil {
		return nil, fmt.Errorf("Subscription.Create: %v", err)
	}
	if scannedEndDate.Valid {
		sub.EndDate = scannedEndDate.String
	} else {
		sub.EndDate = ""
	}

	return sub, nil
}
