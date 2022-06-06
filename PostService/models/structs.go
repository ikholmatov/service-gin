package models

type Address struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	District   string `json:"district"`
	PostalCode int64  `json:"postal_code"`
}
