package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"

	"github.com/go-faker/faker/v4"
	"github.com/ilmsg/hardened-potential/api"
	"github.com/ilmsg/hardened-potential/database"
	"github.com/ilmsg/hardened-potential/model"
)

func main() {
	db := database.NewDatabaseWithSqlite("hardened-potential.db")
	db.AutoMigrate(model.User{}, model.Article{}, model.Book{}, model.Chapter{})

	hChapter := api.NewChapterHandler(db)
	hBook := api.NewBookHandler(db)

	books, _ := hBook.FindAllBook()
	book := books[rand.N(len(books))]

	title, _ := faker.GetLorem().Sentence(reflect.Value{})
	fmt.Printf("%s\n", title)

	chapter, _ := hChapter.NewChapter(title.(string), book.ID)
	fmt.Printf("%s. %s - %s\n", chapter.ID, chapter.Title, chapter.BookId)

	chapters, _ := hChapter.FindAllChapter(book.ID)
	for _, chapter := range chapters {
		fmt.Printf("%s. %s - %s\n", chapter.ID, chapter.Title, chapter.BookId)
	}
}
