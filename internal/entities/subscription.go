package entities

type Subscription struct {
	Id          int    `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserUUID    string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type SubscriptionsFilters struct {
	ServiceName string `json:"service_name" form:"service_name" query:"service_name"`
	UserUUID    string `json:"user_id" form:"user_id" query:"user_id"`
	StartDate   string `json:"start_date" form:"start_date" query:"start_date"`
	EndDate     string `json:"end_date" form:"end_date" query:"end_date"`
}

type SubscriptionAmount struct {
	Amount int `json:"amount"`
}
