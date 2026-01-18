package main

import (
	"fmt"

	"github.com/ilmsg/hardened-potential/api"
	"github.com/ilmsg/hardened-potential/database"
	"github.com/ilmsg/hardened-potential/model"
)

type Person struct {
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

func main() {
	db := database.NewDatabaseWithSqlite("hardened-potential.db")
	db.AutoMigrate(model.User{}, model.Article{}, model.Book{}, model.Chapter{})

	hUser := api.NewUserHandler(db)

	//faker.FakeData(&Person{})
	// hUser.NewUser(fmt.Sprintf("%s %s", person.FirstName, person.LastName))

	persons, _ := hUser.FindAll()
	for _, person := range persons {
		fmt.Printf("%s\n", person.Name)
	}
}
