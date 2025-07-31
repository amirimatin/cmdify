package cmdify

import (
	"fmt"
	"github.com/amirimatin/cmdify/prompt"
	"os"
)

type ConfirmBuilder struct {
	label string
}

func (c *ConfirmBuilder) MustConfirm() {
	ok, err := prompt.AskYesNo(c.label)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Confirmation failed:", err)
		os.Exit(1)
	}
	if !ok {
		fmt.Println("Cancelled.")
		os.Exit(1)
	}
}

func (c *ConfirmBuilder) MustConfirmWithDefaultYes(defaultYes bool) {
	ok, err := prompt.AskYesNoDefault(c.label, defaultYes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Confirmation failed:", err)
		os.Exit(1)
	}
	if !ok {
		fmt.Println("Cancelled.")
		os.Exit(1)
	}
}
