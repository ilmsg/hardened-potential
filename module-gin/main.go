//go:build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	hBook := NewBookHandler()
	hArticle := NewArticleHandler()

	r := gin.Default()

	rBook := r.Group("/books")
	rBook.GET("", hBook.FindAllBook)
	rBook.GET("", hBook.FindBook)
	rBook.GET("", hBook.FindChapter)

	rArticle := r.Group("/article")
	rArticle.GET("", hArticle.FindAllArticle)
	rArticle.GET("", hArticle.FindArticle)

	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}

type BookHandler struct {
}

func (h *BookHandler) FindAllBook(c *gin.Context) {}
func (h *BookHandler) FindBook(c *gin.Context)    {}
func (h *BookHandler) FindChapter(c *gin.Context) {}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

type ArticleHandler struct{}

func (h *ArticleHandler) FindAllArticle(c *gin.Context) {}
func (h *ArticleHandler) FindArticle(c *gin.Context)    {}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}
