package main

import (
	"fmt"
	"github.com/amirimatin/cmdify"
	"github.com/amirimatin/cmdify/prompt"
)

func main() {
	tasks := cmdify.MultiSelect("Select tasks").
		FromList([]prompt.SelectOption{
			{"Build", "Compile the code"},
			{"Test", "Run unit tests"},
			{"Deploy", "Push to production"},
		}).
		Run()

	fmt.Println("Selected tasks:", tasks)
}
