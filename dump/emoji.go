package main

import (
	"fmt"
	"reflect"

	"github.com/go-faker/faker/v4"
	"github.com/ilmsg/hardened-potential/helper"
)

func main() {
	title, _ := faker.GetLorem().Sentence(reflect.Value{})
	// fmt.Printf("%s\n", title)

	emoji := helper.PickRandomEmoji()
	fmt.Printf("%s %s\n", emoji, title)
}
