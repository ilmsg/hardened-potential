//go:build ignore

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ArticleHandler struct{}

func (h *ArticleHandler) FindAllArticle(w http.ResponseWriter, r *http.Request) {}
func (h *ArticleHandler) FindArticle(w http.ResponseWriter, r *http.Request)    {}
func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

type BookHandler struct{}

func (h *BookHandler) FindAllBook(w http.ResponseWriter, r *http.Request) {}
func (h *BookHandler) FindBook(w http.ResponseWriter, r *http.Request)    {}
func (h *BookHandler) FindChapter(w http.ResponseWriter, r *http.Request) {}
func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func main() {
	hArtilce := NewArticleHandler()
	hBook := NewBookHandler()

	r := mux.NewRouter()

	rBook := r.PathPrefix("books").Subrouter()
	rBook.HandleFunc("", hBook.FindAllBook)
	rBook.HandleFunc("/{book_id}", hBook.FindBook)
	rBook.HandleFunc("/{book_id}/chapters/{chapter_id}", hBook.FindChapter)

	rArticle := r.PathPrefix("articles").Subrouter()
	rArticle.HandleFunc("", hArtilce.FindAllArticle)
	rArticle.HandleFunc("/{article_id}", hArtilce.FindArticle)

	log.Fatal(http.ListenAndServe(":7070", r))
}
