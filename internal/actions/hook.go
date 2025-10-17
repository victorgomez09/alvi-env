package actions

import "fmt"

func PrintShellHook() {
	fmt.Println(`
__alenv_hook() {
  # 1. OPTIMIZACIÓN DE RENDIMIENTO: Solo ejecutar la lógica pesada si el directorio cambió.
  if [ "$PWD" != "$__MY_DIRENV_LAST_DIR" ]; then
    
    # 2. Obtiene los comandos export del CLI (usando el alias si está disponible, sino el binario)
    local commands=$(alenv run)

    # 3. Ejecuta los comandos en el shell actual
    eval "$commands"
  fi
  
  # 4. Actualiza el último directorio (variable de estado)
  __MY_DIRENV_LAST_DIR="$PWD"
}

# Variable de estado para la optimización. Se inicializa con el directorio actual.
export __MY_DIRENV_LAST_DIR="$PWD"

# Ejecuta el hook antes de cada prompt
PROMPT_COMMAND="__alenv_hook;$PROMPT_COMMAND"
`)
}
