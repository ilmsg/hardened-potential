package api

type Book struct{}

func (book *Book) FindAllBook() {}
func (book *Book) FindBook()    {}
func (book *Book) FindChapter() {}

func NewBook() *Book {
	return &Book{}
}
