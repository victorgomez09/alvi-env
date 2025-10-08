package actions

import (
	"log"
	"os"

	"github.com/victorgomez09/alvi-env/internal/utils"
	"gopkg.in/yaml.v3"
)

type Config map[string]map[string]interface{}

func InitializeConfig() {
	initialConfig := Config{
		"development": {
			"KEY": "value",
		},
		"testing": {
			"KEY": "value",
		},
		"production": {
			"KEY": "value",
		},
	}

	// Marshall content into YAML format
	yamlData, err := yaml.Marshal(initialConfig)
	if err != nil {
		log.Fatalf("Error marshalling to YAML: %s", err)
	}

	// Write the byte slice to file
	filename := "alenv.yaml"
	if utils.FileExists(filename) {
		log.Println("Alenv config file already exists")
		os.Exit(0)
	}

	err = os.WriteFile(filename, yamlData, 0644)
	if err != nil {
		log.Fatalf("Error writing YAML to file %s: %v", filename, err)
	}

	log.Println("Alenv config file successfully created")
}
