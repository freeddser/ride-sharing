package models

type Passenger struct {
	CustomerID     int64  `json:"customer_id"`
	CustomerName   string `json:"customer_name"`
	CustomerMobile string `json:"customer_mobile"`
	Status         int    `json:"status"`
	CreateAt       string `json:"create_at"`
	UpdateAt       string `json:"update_at"`
}

type GeneratedPassengerResponse struct {
	Success bool `json:"success"`
	Error   struct {
		ErrorID int `json:"error_id"`
		Message struct {
			Cn string `json:"cn"`
			En string `json:"en"`
		} `json:"message"`
	} `json:"error"`
	Results Passenger `json:"results"`
}
