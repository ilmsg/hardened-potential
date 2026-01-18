package api

import (
	"fmt"

	"github.com/ilmsg/hardened-potential/helper"
	"github.com/ilmsg/hardened-potential/model"
	"github.com/ilmsg/hardened-potential/slug"
	"gorm.io/gorm"
)

type IChapterHandler interface {
	NewChapter(title string, bookId string) (*model.Chapter, error)
	FindAllChapter(bookId string) ([]model.Chapter, error)
	FindChapter(chapterId string) (*model.Chapter, error)
}

type ChapterHandler struct {
	db *gorm.DB
}

func (h *ChapterHandler) FindAllChapter(bookId string) ([]model.Chapter, error) {
	var chapters []model.Chapter
	if tx := h.db.Where("book_id = ?", bookId).Find(&chapters); tx.Error != nil {
		return nil, tx.Error
	}
	return chapters, nil
}

func (h *ChapterHandler) FindChapter(chapterId string) (*model.Chapter, error) {
	var chapter model.Chapter
	if tx := h.db.First(&chapter, "id = ?", chapterId); tx.Error != nil {
		return nil, tx.Error
	}
	return &chapter, nil
}

func (h *ChapterHandler) NewChapter(title string, bookId string) (*model.Chapter, error) {
	newChapter := model.Chapter{
		ID:     slug.GenerateSlug(),
		Title:  fmt.Sprintf("%s %s", helper.PickRandomEmoji(), title),
		BookId: bookId,
	}
	if tx := h.db.Create(newChapter); tx.Error != nil {
		return nil, tx.Error
	}

	var chapter model.Chapter
	if tx := h.db.First(&chapter, "id = ?", newChapter.ID); tx.Error != nil {
		return nil, tx.Error
	}
	return &chapter, nil
}

func NewChapterHandler(db *gorm.DB) IChapterHandler {
	return &ChapterHandler{db}
}
