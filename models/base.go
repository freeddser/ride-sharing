package models

type GeneratedDriverResponse struct {
	Success bool `json:"success"`
	Error   struct {
		ErrorID int `json:"error_id"`
		Message struct {
			Cn string `json:"cn"`
			En string `json:"en"`
		} `json:"message"`
	} `json:"error"`
	Results Driver `json:"results"`
}

type GeneratedResponse struct {
	Success bool `json:"success"`
	Error   struct {
		ErrorID int `json:"error_id"`
		Message struct {
			Cn string `json:"cn"`
			En string `json:"en"`
		} `json:"message"`
	} `json:"error"`
}
