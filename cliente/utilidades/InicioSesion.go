package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// Usuario representa un usuario simple en memoria.
type Usuario struct {
	Id       int
	Nickname string
	// La contraseña se almacena en claro solo para fines de práctica;
	// en producción, debe ser un hash seguro.
	Password string
}

// Mapa de usuarios (simulado en una slice global). Cambia o agrega según necesites.
var usuarios = []Usuario{
	{Id: 1, Nickname: "David", Password: "1234"},
	{Id: 2, Nickname: "Camilo", Password: "1234"},
}

// IniciarSesion muestra prompts para solicitar el nickname y la contraseña.
// Devuelve el nickname y el id del usuario autenticado.
// Si la autenticación falla, devuelve ("", 0).
func IniciarSesion() (string, int) {
	reader := bufio.NewReader(os.Stdin)

	// --- 1. Leer Nickname ---
	fmt.Print("Nickname: ")
	nickRaw, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error leyendo nickname:", err)
		return "", 0
	}
	// Eliminar espacios y saltos de línea del input
	nick := strings.TrimSpace(nickRaw)

	// --- 2. Leer Contraseña de forma segura (oculta) ---
	fmt.Print("Contraseña: ")

	// Intenta leer la contraseña sin mostrarla en pantalla
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd())) //os.Stdin.Fd() obtiene el descriptor de archivo
	fmt.Println("")                                            // Asegurar una nueva línea después de la entrada oculta

	if err != nil {
		// En caso de error muestra una advertencia
		// y realiza una lectura visible normal como fallback.
		fmt.Print("\n(Advertencia) No se pudo ocultar la contraseña, se leerá visible.\nContraseña: ")
		passRaw, _ := reader.ReadString('\n')
		bytePassword = []byte(strings.TrimSpace(passRaw))
	}

	// Convertir el slice de bytes a string y limpiar espacios
	password := strings.TrimSpace(string(bytePassword))

	// --- 3. Validar Autenticación ---
	// Iterar sobre el slice de usuarios simulado
	for _, u := range usuarios {
		if u.Nickname == nick && u.Password == password {
			// Éxito: usuario y contraseña coinciden
			fmt.Printf("Bienvenido %s (id=%d)\n", u.Nickname, u.Id)
			return u.Nickname, u.Id
		}
	}

	// Falla: no se encontró ninguna coincidencia después de revisar todos los usuarios
	fmt.Println("El usuario o la contraseña son incorrectos.")
	return "", 0
}
