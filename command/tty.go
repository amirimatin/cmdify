// Package command provides TTY-enabled command execution with proper terminal handling.
package command

import (
	"fmt"
	"github.com/creack/pty"
	"golang.org/x/term"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// ExecuteWithTTY runs a command with full TTY support.
func ExecuteWithTTY(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	ptmx, err := pty.Start(cmd)
	if err != nil {
		return fmt.Errorf("failed to start pty: %w", err)
	}
	defer ptmx.Close()

	// Resize PTY initially
	if err := resizePTY(ptmx); err != nil {
		return fmt.Errorf("failed to resize pty: %w", err)
	}

	// Watch for window resize signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			_ = resizePTY(ptmx)
		}
	}()
	_ = resizePTY(ptmx) // Initial resize

	// Set terminal to raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to set raw mode: %w", err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()

	// Forward I/O
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	return cmd.Wait()
}

// resizePTY resizes the PTY to match the current terminal size.
func resizePTY(ptmx *os.File) error {
	fd := int(os.Stdout.Fd())
	if !term.IsTerminal(fd) {
		return nil
	}
	width, height, err := term.GetSize(fd)
	if err != nil {
		return err
	}
	return pty.Setsize(ptmx, &pty.Winsize{
		Cols: uint16(width),
		Rows: uint16(height),
	})
}
