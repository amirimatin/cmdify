package main

import (
	"fmt"
	"github.com/amirimatin/cmdify"
	"github.com/amirimatin/cmdify/prompt"
)

func main() {
	name := cmdify.Ask("Enter your name").Input()

	cmdify.Ask("Are you sure you want to continue?").YesNo().MustConfirm()

	env := cmdify.Select("Environment").
		FromList([]prompt.SelectOption{
			{"Dev", "Local development"},
			{"Prod", "Live system"},
		}).
		WithDefault(0).
		Run()

	actions := cmdify.MultiSelect("Choose operations").
		FromList([]prompt.SelectOption{
			{"Build", "Compile code"},
			{"Test", "Run tests"},
			{"Deploy", "Deploy to server"},
		}).
		Run()

	fmt.Println("User:", name)
	fmt.Println("Env:", env)
	fmt.Println("Actions:", actions)

	cmdify.Run("echo", "Deployment started...").Must()
}
