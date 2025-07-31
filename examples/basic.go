package main

import "github.com/amirimatin/cmdify"

func main() {
	name := cmdify.Ask("What's your name?").Input()
	password := cmdify.Ask("Enter your password").Password()
	cmdify.Ask("Do you confirm?").YesNo()

	println("Hello", name)
	println("Password:", password)
}
