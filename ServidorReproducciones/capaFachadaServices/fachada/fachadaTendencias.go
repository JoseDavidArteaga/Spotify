package fachada

import (
	"fmt"
	entities "tendencias/capaAccesoDatos/entities"
	"tendencias/capaAccesoDatos/repositorios"
)

type FachadaTendencias struct {
	repo *repositorios.RepositorioReproducciones
}

// NuevaFachadaTendencias es el constructor de la fachada.
func NuevaFachadaTendencias() *FachadaTendencias {
	fmt.Println("Inicializando Fachada de Tendencias...")
	return &FachadaTendencias{
		repo: repositorios.GetRepositorio(),
	}
}

// RegistrarReproduccion registra una nueva reproducción delegando la
// responsabilidad al repositorio.
func (f *FachadaTendencias) RegistrarReproduccion(idUsuario int, titulo string) {
	fmt.Printf("Fachada: Peticion para registrar reproduccion del usuario %d, cancion '%s'\n", idUsuario, titulo)
	f.repo.AgregarReproduccion(idUsuario, titulo)
}

// ObtenerTodasLasReproducciones obtiene todas las reproducciones almacenadas.
func (f *FachadaTendencias) ObtenerTodasLasReproducciones() []entities.ReproduccionEntity {
	// ECO: Imprime que la petición ha llegado a la fachada.
	fmt.Println("Fachada: Peticion para obtener todas las reproducciones.")
	return f.repo.ListarTodasLasReproducciones()
}

// ObtenerReproduccionesPorUsuario obtiene las reproducciones de un usuario específico.
func (f *FachadaTendencias) ObtenerReproduccionesPorUsuario(idUsuario int) []entities.ReproduccionEntity {
	// ECO: Imprime que la petición ha llegado a la fachada.
	fmt.Printf("Fachada: Peticion para obtener reproducciones del usuario %d\n", idUsuario)
	return f.repo.ListarReproduccionesPorUsuario(idUsuario)
}
