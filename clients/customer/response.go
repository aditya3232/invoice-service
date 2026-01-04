package clients

import "time"

type CustomerResponse struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    CustomerData `json:"data"`
}

type CustomerData struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
}
