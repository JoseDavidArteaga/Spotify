package controladores

import (
	"encoding/json"
	"net/http"

	fachada "servidor.local/grpc-servidorCancion/dominio/cancion/fachadaCancionesServices"
)

// ListarCancionesREST maneja las solicitudes HTTP para listar canciones
func ListarCancionesREST(w http.ResponseWriter, r *http.Request) {
	// Verificar que el método sea GET
	if r.Method != http.MethodGet {
		// Responder con error si no es GET
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	canciones := fachada.ObtenerCancionesParaREST()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(canciones)
}
