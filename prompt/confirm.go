package prompt

import "github.com/manifoldco/promptui"

// AskYesNo shows a Yes/No prompt and returns true/false accordingly.
func AskYesNo(question string) (bool, error) {
	prompt := promptui.Select{
		Label: question,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return result == "Yes", nil
}

// AskYesNoDefault shows a Yes/No prompt with default selection (0=Yes, 1=No).
func AskYesNoDefault(question string, defaultYes bool) (bool, error) {
	defaultIndex := 1
	if defaultYes {
		defaultIndex = 0
	}
	prompt := promptui.Select{
		Label:     question,
		Items:     []string{"Yes", "No"},
		CursorPos: defaultIndex,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return result == "Yes", nil
}
