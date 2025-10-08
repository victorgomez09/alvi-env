package utils

import (
	"errors"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	if err == nil {
		return !info.IsDir()
	}

	return false
}
