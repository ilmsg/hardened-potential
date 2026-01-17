package api

import (
	"github.com/ilmsg/hardened-potential/model"
	"gorm.io/gorm"
)

type IBookHandler interface {
	FindAllBook() (*[]model.Book, error)
	FindBook(bookId string) (*model.Book, error)
	FindChapter(chapterId string) (*model.Chapter, error)
}

type BookHandler struct {
	db *gorm.DB
}

func (h *BookHandler) FindAllBook() (*[]model.Book, error) {
	var books []model.Book
	if tx := h.db.Find(&books); tx.Error != nil {
		return nil, tx.Error
	}
	return &books, nil
}

func (h *BookHandler) FindBook(bookId string) (*model.Book, error) {
	var book model.Book
	if tx := h.db.First(&book, "id = ?", bookId); tx.Error != nil {
		return nil, tx.Error
	}
	return &book, nil
}

func (h *BookHandler) FindChapter(chapterId string) (*model.Chapter, error) {
	var chapter model.Chapter
	if tx := h.db.First(&chapter, "id = ?", chapterId); tx.Error != nil {
		return nil, tx.Error
	}
	return &chapter, nil
}

func NewBookHandler(db *gorm.DB) IBookHandler {
	return &BookHandler{db}
}
