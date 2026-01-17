package api

import (
	"github.com/ilmsg/hardened-potential/model"
	"github.com/ilmsg/hardened-potential/slug"
	"gorm.io/gorm"
)

type IArticleHandler interface {
	NewArticle(title string) (*model.Article, error)
	FindAllArticle() (*[]model.Article, error)
	FindArticle(articleId string) (*model.Article, error)
}

type ArticleHandler struct {
	db *gorm.DB
}

func (h *ArticleHandler) NewArticle(title string) (*model.Article, error) {
	newArticle := model.Article{
		ID:    slug.GenerateSlug(),
		Title: title,
	}
	if tx := h.db.Create(&newArticle); tx.Error != nil {
		return nil, tx.Error
	}

	var article model.Article
	if tx := h.db.First(&article, "id = ?", newArticle.ID); tx.Error != nil {
		return nil, tx.Error
	}
	return &article, nil
}

func (h *ArticleHandler) FindAllArticle() (*[]model.Article, error) {
	var articles []model.Article
	if tx := h.db.Find(&articles); tx.Error != nil {
		return nil, tx.Error
	}
	return &articles, nil
}

func (h *ArticleHandler) FindArticle(articleId string) (*model.Article, error) {
	var article model.Article
	if tx := h.db.First(&article, "id = ?", articleId); tx.Error != nil {
		return nil, tx.Error
	}
	return &article, nil
}

func NewArticleHandler(db *gorm.DB) IArticleHandler {
	return &ArticleHandler{db}
}
