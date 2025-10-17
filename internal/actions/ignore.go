package actions

import (
	"fmt"
	"os"
)

const disabledFile = ".alenv-disabled"

func IgnoreEnvrc() error {
	if _, err := os.Stat(envrcFile); os.IsNotExist(err) {
		return fmt.Errorf(".envrc file not exists")
	}

	file, err := os.Create(disabledFile)
	if err != nil {
		return fmt.Errorf("error creating disabled file")
	}

	file.Close()

	return nil
}
