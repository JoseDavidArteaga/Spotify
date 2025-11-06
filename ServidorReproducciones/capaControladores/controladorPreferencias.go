package controladores

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	dtos "tendencias/capaFachadaServices/DTOs"
	"tendencias/capaFachadaServices/fachada"
)

type ControladorTendencias struct {
	fachada *fachada.FachadaTendencias
}

func NuevoControladorTendencias() *ControladorTendencias {
	return &ControladorTendencias{
		fachada: fachada.NuevaFachadaTendencias(),
	}
}

// RegistrarReproduccionHandler gestiona las peticiones POST al endpoint /reproducciones.
func (c *ControladorTendencias) RegistrarReproduccionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var dto dtos.ReproduccionDTOInput
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Cuerpo de la peticion JSON invalido", http.StatusBadRequest)
		return
	}

	c.fachada.RegistrarReproduccion(dto.IdUsuario, dto.Titulo)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Reproduccion registrada correctamente")
}

// ListarReproduccionesHandler gestiona las peticiones GET al endpoint /reproducciones.
func (c *ControladorTendencias) ListarReproduccionesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	idUsuarioStr := r.URL.Query().Get("idUsuario")

	var reproducciones interface{}

	if idUsuarioStr != "" {
		idUsuario, err := strconv.Atoi(idUsuarioStr)
		if err != nil {
			http.Error(w, "El parametro 'idUsuario' debe ser un numero entero.", http.StatusBadRequest)
			return
		}
		reproducciones = c.fachada.ObtenerReproduccionesPorUsuario(idUsuario)
	} else {
		reproducciones = c.fachada.ObtenerTodasLasReproducciones()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reproducciones)
}
