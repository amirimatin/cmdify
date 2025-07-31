// Package command provides utilities for running shell commands with or without TTY.
package command

import (
	"bytes"
	"context"
	"os/exec"
)

// CommandResult holds the result of a command execution.
type CommandResult struct {
	Output   string
	ExitCode int
	Err      error
}

// Execute runs a command and returns its output, exit code, and any error.
func Execute(ctx context.Context, name string, args ...string) CommandResult {
	cmd := exec.CommandContext(ctx, name, args...)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = 1
		}
	}
	return CommandResult{
		Output:   out.String(),
		ExitCode: exitCode,
		Err:      err,
	}
}
