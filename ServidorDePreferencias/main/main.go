package main

import (
	"fmt"
	"log"
	"net/http"
	"preferencias/controladores"
)

func main() {
	fmt.Println("   Servidor de Preferencias Musicales")
	puerto := ":2021"

	// Crear el controlador principal
	ctrl := controladores.NuevoControladorPreferencias()

	// Registrar el endpoint para calcular preferencias
	http.HandleFunc("/preferencias/calcular", ctrl.CalcularPreferenciasHandler)

	fmt.Println("Servidor iniciado correctamente")
	fmt.Printf("Escuchando en el puerto %s\n", puerto)

	// Iniciar el servidor HTTP
	if err := http.ListenAndServe(puerto, nil); err != nil {
		log.Fatalf("Error crítico al iniciar el servidor: %v", err)
	}
}
