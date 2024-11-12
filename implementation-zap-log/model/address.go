package model

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	PostalCode int    `json:"postal_code"`
	Country    string `json:"country"`
}
