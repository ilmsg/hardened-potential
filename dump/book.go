package main

import (
	"fmt"

	"github.com/ilmsg/hardened-potential/api"
	"github.com/ilmsg/hardened-potential/database"
	"github.com/ilmsg/hardened-potential/model"
)

func main() {
	db := database.NewDatabaseWithSqlite("hardened-potential.db")
	db.AutoMigrate(model.User{}, model.Article{}, model.Book{}, model.Chapter{})

	hBook := api.NewBookHandler(db)
	hChapter := api.NewChapterHandler(db)
	// hUser := api.NewUserHandler(db)

	// users, _ := hUser.FindAll()
	// user := users[rand.N(len(users))]

	// title, _ := faker.GetLorem().Sentence(reflect.Value{})
	// fmt.Printf("%s\n", title)

	// book, _ := hBook.NewBook(title.(string), user.ID)
	// fmt.Printf("%s. %s - %s\n", book.ID, book.Title, book.UserId)

	books, _ := hBook.FindAllBook()
	for _, book := range books {
		fmt.Printf("-----------------------------------------\n")
		fmt.Printf("%s. %s - %s\n", book.ID, book.Title, book.UserId)
		chapters, _ := hChapter.FindAllChapter(book.ID)
		for _, chapter := range chapters {
			fmt.Printf("%s. %s\n", chapter.ID, chapter.Title)
		}
	}
}
