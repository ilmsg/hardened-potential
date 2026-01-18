package model

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

type Chapter struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	BookId string `json:"book_id"`
}
