package model

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Chapter struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	BookId uint   `json:"book_id"`
}
