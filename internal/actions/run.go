package actions

import (
	"fmt"
	"log"
	"os"

	"github.com/victorgomez09/alvi-env/internal/utils"
	"gopkg.in/yaml.v3"
)

func RunEnvironment(environment string) {
	filename := "alenv.yaml"

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading Alenv configuration file %s", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML from file: %v", err)
	}

	envVars, ok := config[environment]
	if !ok {
		log.Fatalf("Environment key '%s' not found in YAML.", environment)
	}

	fmt.Printf("--- Reading variables for environment: %s from file: %s ---\n", environment, filename)

	// 4. Iterate and print all key/value pairs for that environment
	for key, value := range envVars {
		envValue := fmt.Sprintf("%v", value)
		err := os.Setenv(key, envValue)
		if err != nil {
			log.Printf("Warning: Failed to set environment variable %s: %v", key, err)
			continue
		}

		utils.SetEnvironmentVariables(key, envValue)
		// fmt.Printf("  ✅ Set %s=%s\n", key, envValue)
	}
}
