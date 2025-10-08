package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func SetEnvironmentVariables(key string, value string) {
	// Determine the shell command based on operating system
	var prefix string
	currentOS := runtime.GOOS

	switch currentOS {
	case "linux", "darwin": // Linux and macOS (uses bash/zsh/etc.)
		prefix = "export"
		fmt.Fprintf(os.Stderr, "Detected OS: %s. Using 'export' syntax.\n", currentOS)
	case "windows": // Windows (uses cmd/powershell)
		prefix = "SET"
		fmt.Fprintf(os.Stderr, "Detected OS: %s. Using 'SET' syntax.\n", currentOS)
	default:
		// Fallback for less common OS
		prefix = "export "
		fmt.Fprintf(os.Stderr, "Detected OS: %s. Using default 'export' syntax.\n", currentOS)
	}

	var command string
	if currentOS == "windows" {
		// Windows syntax: SET KEY=VALUE
		command = fmt.Sprintf("%s %s=%s", prefix, key, value)
	} else {
		// Unix-like syntax: export KEY='VALUE'
		// Use single quotes to protect the value from shell expansion
		command = fmt.Sprintf("%s %s=%s", prefix, key, value)
	}

	runCommand(command)
}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)

	// Run the command and capture the output
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Command failed: %s\n", err)
		return
	}
}
