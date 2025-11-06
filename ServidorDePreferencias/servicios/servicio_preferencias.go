package servicios

import (
	"fmt"
	"preferencias/clientes"
	"preferencias/modelos"
)

// ServicioPreferencias orquesta la colaboración entre diferentes componentes
// para calcular las preferencias de un usuario
type ServicioPreferencias struct {
	clienteCanciones       *clientes.ClienteCanciones
	clienteReproducciones  *clientes.ClienteReproducciones
	calculadorPreferencias *CalculadorPreferencias
}

// NuevoServicioPreferencias crea una nueva instancia del servicio con todas sus dependencias
func NuevoServicioPreferencias() *ServicioPreferencias {
	return &ServicioPreferencias{
		clienteCanciones:       clientes.NuevoClienteCanciones(),
		clienteReproducciones:  clientes.NuevoClienteReproducciones(),
		calculadorPreferencias: NuevoCalculadorPreferencias(),
	}
}

// ObtenerPreferencias calcula las preferencias musicales de un usuario
// Este método orquesta el flujo de trabajo completo:
// 1. Obtiene el catálogo completo de canciones
// 2. Obtiene el historial de reproducciones del usuario
// 3. Calcula y retorna las preferencias agregadas
func (s *ServicioPreferencias) ObtenerPreferencias(idUsuario int) (modelos.PreferenciasRespuesta, error) {
	fmt.Printf("--> ServicioPreferencias: Obteniendo datos para el usuario con ID: %d\n", idUsuario)

	// 1. Obtener el catálogo completo de canciones
	catalogoCanciones, err := s.clienteCanciones.ObtenerCanciones()
	if err != nil {
		return modelos.PreferenciasRespuesta{}, fmt.Errorf("error al obtener canciones: %v", err)
	}

	// ECO: Imprimir las canciones obtenidas
	fmt.Println("    ServicioPreferencias: Canciones obtenidas del Servidor de Canciones:")
	for _, cancion := range catalogoCanciones {
		fmt.Printf("      - Titulo: %s, Artista: %s, Genero: %s, Idioma: %s\n",
			cancion.Titulo, cancion.Artista, cancion.Genero, cancion.Idioma)
	}

	// 2. Obtener el historial de reproducciones del usuario
	reproduccionesUsuario, err := s.clienteReproducciones.ObtenerReproducciones(idUsuario)
	if err != nil {
		return modelos.PreferenciasRespuesta{}, fmt.Errorf("error al obtener reproducciones: %v", err)
	}

	// ECO: Imprimir las reproducciones obtenidas
	fmt.Printf("    ServicioPreferencias: Reproducciones obtenidas del Servidor de Reproducciones para el usuario %d:\n", idUsuario)
	for _, reproduccion := range reproduccionesUsuario {
		fmt.Printf("      - Titulo: %s, Fecha: %s\n",
			reproduccion.Titulo, reproduccion.FechaHora)
	}

	// 3. Calcular las preferencias usando el calculador
	preferencias := s.calculadorPreferencias.Calcular(idUsuario, catalogoCanciones, reproduccionesUsuario)

	return preferencias, nil
}
