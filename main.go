package main

import (
	"fmt"

	"github.com/ilmsg/hardened-potential/slug"
)

func main() {
	newSlug := slug.GenerateSlug()
	fmt.Printf("Generate Slug: %s\n", newSlug)
}
