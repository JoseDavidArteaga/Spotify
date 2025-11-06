package clientes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"preferencias/modelos"
)

// ClienteCanciones es responsable de la comunicación con el Servidor de Canciones
type ClienteCanciones struct {
	baseURL string
}

// NuevoClienteCanciones crea una nueva instancia del cliente
func NuevoClienteCanciones() *ClienteCanciones {
	return &ClienteCanciones{
		baseURL: "http://localhost:5000",
	}
}

// ObtenerCanciones realiza una petición GET al Servidor de Canciones para obtener el catálogo completo
func (c *ClienteCanciones) ObtenerCanciones() ([]modelos.CancionDTO, error) {
	url := c.baseURL + "/canciones"
	fmt.Println("ClienteCanciones: Realizando petición GET a", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al comunicar con el Servidor de Canciones: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("el Servidor de Canciones respondió con código %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer respuesta del Servidor de Canciones: %v", err)
	}

	var canciones []modelos.CancionDTO
	if err := json.Unmarshal(body, &canciones); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON de canciones: %v", err)
	}

	fmt.Printf("ClienteCanciones: Se recibieron %d canciones\n", len(canciones))
	return canciones, nil
}
