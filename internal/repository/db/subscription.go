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

func (dbr *DatabaseRepository) GetSumSubscriptionByFilters(filters *entities.SubscriptionsFilters) (int, error) {

	selectQuery := `
		SELECT SUM(price) AS sum_price
		FROM subscription
		WHERE 1=1`

	var sum int
	if filters.ServiceName != "" {
		selectQuery += fmt.Sprintf(" AND name = '%s'", filters.ServiceName)
	}

	if filters.UserUUID != "" {
		selectQuery += fmt.Sprintf(" AND user_uuid = '%s'", filters.UserUUID)
	}

	if filters.StartDate != "" {
		selectQuery += fmt.Sprintf(" AND date_start >= '%s'", filters.StartDate)
	}

	if filters.EndDate != "" {
		selectQuery += fmt.Sprintf(" AND date_end <= '%s'", filters.EndDate)
	}

	err := dbr.DB.QueryRow(selectQuery).Scan(&sum)
	if err != nil {
		return 0, fmt.Errorf("Subscription.GetSumByFilters: %v", err)
	}
	return sum, nil
}

func (dbr *DatabaseRepository) UpdateSubscription(sub *entities.Subscription) (*entities.Subscription, error) {
	var endDate interface{}
	if sub.EndDate == "" {
		endDate = nil
	} else {
		endDate = sub.EndDate
	}

	updateQuery := `
		UPDATE subscription 
		SET name = $1, price = $2, user_uuid = $3, date_start = $4, date_end = $5
		WHERE id = $6
		RETURNING id, name, price, user_uuid, date_start, date_end;
	`

	var scannedEndDate sql.NullString
	err := dbr.DB.QueryRow(updateQuery, sub.ServiceName, sub.Price, sub.UserUUID, sub.StartDate, endDate, sub.Id).Scan(
		&sub.Id, &sub.ServiceName, &sub.Price, &sub.UserUUID, &sub.StartDate, &scannedEndDate)
	if err != nil {
		return nil, fmt.Errorf("Subscription.Update: %v", err)
	}
	if scannedEndDate.Valid {
		sub.EndDate = scannedEndDate.String
	} else {
		sub.EndDate = ""
	}

	return sub, nil
}

func (dbr *DatabaseRepository) DeleteSubscription(id int) error {

	deleteQuery := `
		DELETE FROM subscription WHERE id = $1;
	`

	_, err := dbr.DB.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("Subscription.Delete: %v", err)
	}
	return nil
}
