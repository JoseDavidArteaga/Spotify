package repositorios

import (
	"fmt"
	"strings"
	"sync"
	entities "tendencias/capaAccesoDatos/entities"
	"time"
)

type RepositorioReproducciones struct {
	mu             sync.Mutex
	reproducciones []entities.ReproduccionEntity
}

var (
	instancia *RepositorioReproducciones
	once      sync.Once
)

func GetRepositorio() *RepositorioReproducciones {
	once.Do(func() {
		instancia = &RepositorioReproducciones{}
		instancia.poblarDatosDeEjemplo()
	})
	return instancia
}

func (r *RepositorioReproducciones) poblarDatosDeEjemplo() {
	r.reproducciones = []entities.ReproduccionEntity{
		{IdUsuario: 1, Titulo: "Pedro Navaja", FechaHora: "2025-10-20 10:00:00"},
		{IdUsuario: 1, Titulo: "La Rebelión", FechaHora: "2025-10-20 10:05:00"},
		{IdUsuario: 2, Titulo: "Tu eres Ajena", FechaHora: "2025-10-20 10:10:00"},
		{IdUsuario: 1, Titulo: "La Bilirrubina", FechaHora: "2025-10-20 10:15:00"},
		{IdUsuario: 2, Titulo: "De Música Ligera", FechaHora: "2025-10-20 10:20:00"},
		{IdUsuario: 1, Titulo: "Lamento Boliviano", FechaHora: "2025-10-21 11:00:00"},
		{IdUsuario: 3, Titulo: "Bailando", FechaHora: "2025-10-21 11:05:00"},
		{IdUsuario: 2, Titulo: "Despacito", FechaHora: "2025-10-21 11:10:00"},
		{IdUsuario: 3, Titulo: "Master of Puppets", FechaHora: "2025-10-21 11:15:00"},
		{IdUsuario: 1, Titulo: "The Number of the Beast", FechaHora: "2025-10-21 11:20:00"},
	}
	fmt.Println("Repositorio de Reproducciones inicializado con datos de ejemplo.")
}

// AgregarReproduccion añade un nuevo registro de reproducción a la colección en
func (r *RepositorioReproducciones) AgregarReproduccion(idUsuario int, titulo string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tituloLimpio := strings.TrimSuffix(titulo, ".mp3")

	nuevaReproduccion := entities.ReproduccionEntity{
		IdUsuario: idUsuario,
		Titulo:    tituloLimpio,
		FechaHora: time.Now().Format("2006-01-02 15:04:05"),
	}

	r.reproducciones = append(r.reproducciones, nuevaReproduccion)
	fmt.Printf("Reproduccion almacenada en el repositorio: %+v\n", nuevaReproduccion)
}

// ListarTodasLasReproducciones devuelve un slice con todos los registros de reproduccion
func (r *RepositorioReproducciones) ListarTodasLasReproducciones() []entities.ReproduccionEntity {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.reproducciones
}

// ListarReproduccionesPorUsuario busca y devuelve todos los registros de reproducción
func (r *RepositorioReproducciones) ListarReproduccionesPorUsuario(idUsuario int) []entities.ReproduccionEntity {
	r.mu.Lock()
	defer r.mu.Unlock()

	reproduccionesDelUsuario := []entities.ReproduccionEntity{}
	for _, repro := range r.reproducciones {
		if repro.IdUsuario == idUsuario {
			reproduccionesDelUsuario = append(reproduccionesDelUsuario, repro)
		}
	}

	fmt.Printf("Consulta al repositorio: Se encontraron %d reproducciones para el usuario %d\n", len(reproduccionesDelUsuario), idUsuario)
	return reproduccionesDelUsuario
}
