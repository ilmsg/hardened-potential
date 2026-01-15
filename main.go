package main

import (
	"fmt"

	"github.com/ilmsg/hardened-potential/helper"
	"github.com/ilmsg/hardened-potential/slug"
)

func main() {
	newSlug := slug.GenerateSlug()
	fmt.Printf("Generate Slug: %s\n", newSlug)

	pickEmoji := helper.PickRandomEmoji()
	fmt.Printf("Pick Random Emoji: %s\n", pickEmoji)
}
