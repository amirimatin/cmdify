package cmdify

import (
	"fmt"
	"github.com/amirimatin/cmdify/prompt"
	"os"
)

type SelectBuilder struct {
	label        string
	items        []prompt.SelectOption
	defaultIndex int
}

// Select starts a select prompt builder
func Select(label string) *SelectBuilder {
	return &SelectBuilder{
		label:        label,
		defaultIndex: 0,
	}
}

func (s *SelectBuilder) FromList(options []prompt.SelectOption) *SelectBuilder {
	s.items = options
	return s
}

func (s *SelectBuilder) WithDefault(index int) *SelectBuilder {
	s.defaultIndex = index
	return s
}

func (s *SelectBuilder) Run() string {
	result, err := prompt.AskSelectOptionDefault(s.label, s.items, s.defaultIndex)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Selection failed:", err)
		os.Exit(1)
	}
	return result
}
