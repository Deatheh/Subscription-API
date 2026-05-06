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

	insertQuery := `
		INSERT INTO subscription (name, price, user_uuid, date_start, date_end)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, price, user_uuid, date_start, date_end;
	`

	var scannedEndDate sql.NullString
	err := dbr.DB.QueryRow(insertQuery, sub.ServiceName, sub.Price, sub.UserUUID, sub.StartDate, endDate).Scan(
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

func (dbr *DatabaseRepository) GetAllSubscription() (*[]entities.Subscription, error) {

	selectQuery := `
		SELECT id, name, price, user_uuid, date_start, date_end
		FROM subscription;
	`

	rows, err := dbr.DB.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("Subscription.GetAll: %v", err)
	}
	defer rows.Close()

	var subscriptionSlice []entities.Subscription

	for rows.Next() {
		var sub entities.Subscription
		var scannedEndDate sql.NullString
		err := rows.Scan(&sub.Id, &sub.ServiceName, &sub.Price, &sub.UserUUID, &sub.StartDate, &scannedEndDate)
		if err != nil {
			return nil, fmt.Errorf("Subscription.GetAll: %v", err)
		}
		if scannedEndDate.Valid {
			sub.EndDate = scannedEndDate.String
		} else {
			sub.EndDate = ""
		}
		subscriptionSlice = append(subscriptionSlice, sub)
	}

	return &subscriptionSlice, nil
}

func (dbr *DatabaseRepository) GetSubscriptionById(id int) (*entities.Subscription, error) {

	selectQuery := `
		SELECT id, name, price, user_uuid, date_start, date_end
		FROM subscription
		WHERE id = $1;
	`

	row := dbr.DB.QueryRow(selectQuery, id)

	var sub entities.Subscription
	var scannedEndDate sql.NullString
	err := row.Scan(&sub.Id, &sub.ServiceName, &sub.Price, &sub.UserUUID, &sub.StartDate, &scannedEndDate)
	if err != nil {
		return nil, fmt.Errorf("Subscription.GetById: %v", err)
	}
	if scannedEndDate.Valid {
		sub.EndDate = scannedEndDate.String
	} else {
		sub.EndDate = ""
	}

	return &sub, nil
}
