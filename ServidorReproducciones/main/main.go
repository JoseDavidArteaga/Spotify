package main

import (
	"fmt"
	"log"
	"net/http"
	controladores "tendencias/capaControladores"
)

func main() {
	fmt.Println("Iniciando Servidor de Reproducciones...")

	puerto := ":5002"
	ctrl := controladores.NuevoControladorTendencias()

	// Manejador único para el endpoint /reproducciones que delega según el método HTTP.
	http.HandleFunc("/reproducciones", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ctrl.ListarReproduccionesHandler(w, r)
		case http.MethodPost:
			ctrl.RegistrarReproduccionHandler(w, r)
		default:
			http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf(" Servidor de Reproducciones escuchando en el puerto %s\n", puerto)

	fmt.Println("Servidor iniciado correctamente. Esperando solicitudes...")

	if err := http.ListenAndServe(puerto, nil); err != nil {
		log.Fatalf("Error critico al iniciar el servidor en el puerto %s: %v", puerto, err)
	}
}
