package model

type Article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}
