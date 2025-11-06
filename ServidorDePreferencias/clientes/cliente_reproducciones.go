package clientes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"preferencias/modelos"
)

// ClienteReproducciones es responsable de la comunicación con el Servidor de Reproducciones
type ClienteReproducciones struct {
	baseURL string
}

// NuevoClienteReproducciones crea una nueva instancia del cliente
func NuevoClienteReproducciones() *ClienteReproducciones {
	return &ClienteReproducciones{
		baseURL: "http://localhost:5002",
	}
}

// ObtenerReproducciones realiza una petición GET al Servidor de Reproducciones
// para obtener el historial de un usuario específico
func (c *ClienteReproducciones) ObtenerReproducciones(idUsuario int) ([]modelos.ReproduccionDTO, error) {
	url := fmt.Sprintf("%s/reproducciones?idUsuario=%d", c.baseURL, idUsuario)
	fmt.Println("--> ClienteReproducciones: Realizando petición GET a", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al comunicar con el Servidor de Reproducciones: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("el Servidor de Reproducciones respondió con código %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer respuesta del Servidor de Reproducciones: %v", err)
	}

	var reproducciones []modelos.ReproduccionDTO
	if err := json.Unmarshal(body, &reproducciones); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON de reproducciones: %v", err)
	}

	fmt.Printf("--> ClienteReproducciones: Se recibieron %d reproducciones para el usuario %d\n", len(reproducciones), idUsuario)
	return reproducciones, nil
}
