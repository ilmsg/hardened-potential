package main

import (
	"fmt"

	"github.com/ilmsg/hardened-potential/api"
	"github.com/ilmsg/hardened-potential/database"
	"github.com/ilmsg/hardened-potential/model"
)

type ArticleTitle struct {
}

func main() {
	db := database.NewDatabaseWithSqlite("hardened-potential.db")
	db.AutoMigrate(model.User{}, model.Article{}, model.Book{}, model.Chapter{})

	hArticle := api.NewArticleHandler(db)
	// hUser := api.NewUserHandler(db)

	// users, _ := hUser.FindAll()
	// user := users[rand.N(len(users))]

	// title, _ := faker.GetLorem().Sentence(reflect.Value{})
	// fmt.Printf("%s\n", title)

	// article, _ := hArticle.NewArticle(title.(string), user.ID)
	// fmt.Printf("%s. %s - %s\n", article.ID, article.Title, article.UserId)

	articles, _ := hArticle.FindAllArticle()
	for _, article := range articles {
		fmt.Printf("%s. %s - %s\n", article.ID, article.Title, article.UserId)
	}
}
