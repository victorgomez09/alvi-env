package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const envrcFile = ".envrc"

func ExecuteEnvrc() error {
	if _, err := os.Stat(envrcFile); os.IsNotExist(err) {
		return nil
	}

	if _, err := os.Stat(disabledFile); err == nil {
		return nil
	}

	file, err := os.Open(envrcFile)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo %s: %w", envrcFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pathToAdd string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// **FIX DE PARSING PARA RECARGA:** Si el valor está entre comillas, las quitamos.
		// Esto evita que las comillas dobles terminen dentro de las comillas simples de export.
		if len(value) > 1 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
		}

		// La clave: Usar comillas simples para la exportación y escapar cualquier comilla interna.
		fmt.Printf("export %s='%s'\n", key, strings.ReplaceAll(value, "'", "'\"'\"'"))

		if key == "CLI_BIN_PATH" {
			pathToAdd = value
		}
	}

	if pathToAdd != "" {
		fmt.Printf("export PATH=\"%s:$PATH\"\n", pathToAdd)
	}

	fmt.Fprintf(os.Stderr, "Variables cargadas desde %s\n", envrcFile)

	return scanner.Err()
}

// func RunEnvironment(environment string) {
// 	filename := "alenv.yaml"

// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		log.Fatalf("Error reading Alenv configuration file %s", err)
// 	}

// 	var config Config
// 	err = yaml.Unmarshal(data, &config)
// 	if err != nil {
// 		log.Fatalf("Error unmarshalling YAML from file: %v", err)
// 	}

// 	envVars, ok := config[environment]
// 	if !ok {
// 		log.Fatalf("Environment key '%s' not found in YAML.", environment)
// 	}

// 	fmt.Printf("--- Reading variables for environment: %s from file: %s ---\n", environment, filename)

// 	// 4. Iterate and print all key/value pairs for that environment
// 	for key, value := range envVars {
// 		envValue := fmt.Sprintf("%v", value)
// 		err := os.Setenv(key, envValue)
// 		if err != nil {
// 			log.Printf("Warning: Failed to set environment variable %s: %v", key, err)
// 			continue
// 		}

// 		utils.SetEnvironmentVariables(key, envValue)
// 		// fmt.Printf("  ✅ Set %s=%s\n", key, envValue)
// 	}
// }
