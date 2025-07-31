package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

// SelectOption represents an item in the selection list.
type SelectOption struct {
	Name        string // Displayed title
	Description string // Additional info
}

// AskSelectOption prompts the user to select from a list with optional search.
func AskSelectOption(label string, items []SelectOption) (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "➤ {{ .Name | yellow }} {{ .Description | faint }}",
		Inactive: "  {{ .Name | cyan }} {{ .Description | faint }}",
		Selected: "{{ .Name | green }}",
		Details: `
-------- Detail --------
{{ "Name:" | faint }}        {{ .Name }}
{{ "Description:" | faint }} {{ .Description }}`,
	}

	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.ToLower(strings.ReplaceAll(item.Name, " ", ""))
		input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:             label,
		Items:             items,
		Templates:         templates,
		Size:              8,
		Searcher:          searcher,
		StartInSearchMode: false,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("selection failed: %w", err)
	}

	return items[idx].Name, nil
}

// AskSelectOptionDefault prompts the user to select an option with a default preselected index.
func AskSelectOptionDefault(label string, items []SelectOption, defaultIndex int) (string, error) {
	if defaultIndex < 0 || defaultIndex >= len(items) {
		defaultIndex = 0 // fallback
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "➤ {{ .Name | yellow }} {{ .Description | faint }}",
		Inactive: "  {{ .Name | cyan }} {{ .Description | faint }}",
		Selected: "{{ .Name | green }}",
		Details: `
-------- Detail --------
{{ "Name:" | faint }}        {{ .Name }}
{{ "Description:" | faint }} {{ .Description }}`,
	}

	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.ToLower(strings.ReplaceAll(item.Name, " ", ""))
		input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
		CursorPos: defaultIndex,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("selection failed: %w", err)
	}
	return items[idx].Name, nil
}
