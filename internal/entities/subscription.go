package entities

type Subscription struct {
	Id          string `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserUUID    string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type SubscriptionsFilters struct {
	ServiceName string `json:"service_name"`
	UserUUID    string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type SubscriptionAmount struct {
	Amount int `json:"amount"`
}