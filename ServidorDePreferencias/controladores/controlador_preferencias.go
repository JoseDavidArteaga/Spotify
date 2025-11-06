package controladores

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"preferencias/modelos"
	"preferencias/servicios"
)

// ControladorPreferencias maneja las peticiones HTTP relacionadas con preferencias
type ControladorPreferencias struct {
	servicio *servicios.ServicioPreferencias
}

// NuevoControladorPreferencias crea una nueva instancia del controlador
func NuevoControladorPreferencias() *ControladorPreferencias {
	return &ControladorPreferencias{
		servicio: servicios.NuevoServicioPreferencias(),
	}
}

// CalcularPreferenciasHandler maneja las peticiones POST a /preferencias/calcular
func (ctrl *ControladorPreferencias) CalcularPreferenciasHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que sea una petición POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido. Use POST", http.StatusMethodNotAllowed)
		return
	}

	// Leer el cuerpo de la petición
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la petición", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Decodificar el JSON de la solicitud
	var solicitud modelos.SolicitudPreferencias
	if err := json.Unmarshal(body, &solicitud); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("\n=== Nueva petición de preferencias ===\n")
	fmt.Printf("Usuario ID: %d\n", solicitud.IdUsuario)

	// Validar que el ID de usuario sea válido
	if solicitud.IdUsuario <= 0 {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para obtener las preferencias
	preferencias, err := ctrl.servicio.ObtenerPreferencias(solicitud.IdUsuario)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		http.Error(w, "Error al calcular preferencias: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir la respuesta a JSON
	respuestaJSON, err := json.Marshal(preferencias)
	if err != nil {
		http.Error(w, "Error al codificar respuesta JSON", http.StatusInternalServerError)
		return
	}

	// Enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respuestaJSON)

	fmt.Printf("=== Preferencias calculadas exitosamente ===\n\n")
}
