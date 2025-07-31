package main

import "github.com/amirimatin/cmdify"

func main() {
	cmdify.Run("ls", "-la").Must()
	// Or:
	cmdify.Run("htop").WithTTY().Must()
}
