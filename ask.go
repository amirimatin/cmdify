package cmdify

import (
	"fmt"
	"github.com/amirimatin/cmdify/prompt"
	"os"
)

type AskBuilder struct {
	label string
}

// Ask starts an input prompt chain.
func Ask(label string) *AskBuilder {
	return &AskBuilder{label: label}
}

func (a *AskBuilder) Input() string {
	result, err := prompt.AskInput(a.label)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read input:", err)
		os.Exit(1)
	}
	return result
}

func (a *AskBuilder) Password() string {
	result, err := prompt.AskPassword(a.label)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read password:", err)
		os.Exit(1)
	}
	return result
}

func (a *AskBuilder) YesNo() (bool, error) {
	return prompt.AskYesNo(a.label)
}

func (a *AskBuilder) YesNoDefaultYes(defaultYes bool) (bool, error) {
	return prompt.AskYesNoDefault(a.label, defaultYes)
}

func (a *AskBuilder) DangerConfirm() {
	ok, err := prompt.AskYesNo(a.label + " (this action is irreversible!)")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Confirmation failed:", err)
		os.Exit(1)
	}
	if !ok {
		fmt.Println("Cancelled.")
		os.Exit(1)
	}
}
