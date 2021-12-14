package utils

import (
	"strings"
)

// Splits the given command into the command
// name and the arguments
func SplitCommand(command string) (string, string) {
	splitted := strings.SplitN(command, " ", 2)

	// check for arguments
	cmd := splitted[0]
	args := ""
	if len(splitted) > 1 {
		args = splitted[1]
	}

	return cmd, args
}
