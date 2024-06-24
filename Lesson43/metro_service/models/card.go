package models

type Card struct {
	Id     string `json:"id"`
	Number string `json:"number"`
	UserId string `json:"user_id"`
}

type CreateCard struct {
	Number string `json:"number"`
	UserId string `json:"user_id"`
}
