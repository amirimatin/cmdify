package cmdify

import (
	"fmt"
	"github.com/amirimatin/cmdify/prompt"
	"os"
)

type MultiSelectBuilder struct {
	label string
	items []prompt.SelectOption
}

// MultiSelect starts a multi-selection chain
func MultiSelect(label string) *MultiSelectBuilder {
	return &MultiSelectBuilder{
		label: label,
	}
}

func (m *MultiSelectBuilder) FromList(items []prompt.SelectOption) *MultiSelectBuilder {
	m.items = items
	return m
}

// Run repeatedly prompts user to pick items
func (m *MultiSelectBuilder) Run() []string {
	selected := make([]string, 0)
	remaining := make([]prompt.SelectOption, len(m.items))
	copy(remaining, m.items)

	for len(remaining) > 0 {
		// Add "Done" option
		allOptions := append([]prompt.SelectOption{{Name: "✓ Done", Description: "Finish selection"}}, remaining...)

		choice, err := prompt.AskSelectOptionDefault(m.label+" (select multiple)", allOptions, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Multi-select failed:", err)
			os.Exit(1)
		}
		if choice == "✓ Done" {
			break
		}

		selected = append(selected, choice)

		// Remove selected from remaining
		filtered := make([]prompt.SelectOption, 0)
		for _, item := range remaining {
			if item.Name != choice {
				filtered = append(filtered, item)
			}
		}
		remaining = filtered
	}

	return selected
}
