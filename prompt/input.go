// Package prompt provides interactive CLI input utilities.
package prompt

import (
	"github.com/manifoldco/promptui"
	"io"
	"strings"
)

type writeCloser struct {
	io.Writer
}

func (wc writeCloser) Close() error {
	return nil
}

// AskInput shows a simple input prompt to the user.
func AskInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}
	result, err := prompt.Run()
	return strings.TrimSpace(result), err
}

// AskPassword securely prompts the user to enter a password.
func AskPassword(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Mask:  '*',
	}
	result, err := prompt.Run()
	return strings.TrimSpace(result), err
}

func AskInputWithIO(label string, in io.Reader, out io.Writer) (string, error) {
	prompt := promptui.Prompt{
		Label:  label,
		Stdin:  io.NopCloser(in),
		Stdout: writeCloser{out},
	}
	result, err := prompt.Run()
	return strings.TrimSpace(result), err
}
