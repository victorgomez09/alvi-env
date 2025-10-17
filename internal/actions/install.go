package actions

import "fmt"

func InstallEnvrc() {
	fmt.Println(`
	alenv() {
    if [ "$1" = "reload" ]; then
        # La clave: El shell padre ejecuta el eval para recargar.
        # 'command alenv' llama al binario real y evita la recursión.
        eval "$(command alenv run)"
    else
        # Pasa todos los demás comandos (hook, ignore, run) al binario real.
        # 'command' asegura que se ejecuta el binario y no esta función de nuevo.
        command alenv "$@"
    fi
}
	`)
}
