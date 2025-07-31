package main

import (
	"fmt"
	"github.com/amirimatin/cmdify"
	"github.com/amirimatin/cmdify/prompt"
)

func main() {
	choice := cmdify.Select("Choose environment").
		FromList([]prompt.SelectOption{
			{"Development", "Dev machine"},
			{"Staging", "Pre-prod"},
			{"Production", "Live system"},
		}).
		WithDefault(0).
		Run()

	fmt.Println("Selected:", choice)
}
