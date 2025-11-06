package servicios

import (
	"fmt"
	"preferencias/modelos"
	"sort"
)

// CalculadorPreferencias contiene la lógica de negocio para calcular preferencias musicales
type CalculadorPreferencias struct{}

// NuevoCalculadorPreferencias crea una nueva instancia del calculador
func NuevoCalculadorPreferencias() *CalculadorPreferencias {
	return &CalculadorPreferencias{}
}

// Calcular procesa el catálogo de canciones y el historial de reproducciones
// para generar estadísticas agregadas por género, artista e idioma
func (calc *CalculadorPreferencias) Calcular(
	idUsuario int,
	catalogoCanciones []modelos.CancionDTO,
	reproduccionesUsuario []modelos.ReproduccionDTO,
) modelos.PreferenciasRespuesta {

	fmt.Printf("--> CalculadorPreferencias: Iniciando cálculo para el usuario %d\n", idUsuario)
	fmt.Printf("    Recibidas %d canciones del catálogo y %d reproducciones\n",
		len(catalogoCanciones), len(reproduccionesUsuario))

	mapaCatalogo := make(map[string]modelos.CancionDTO)
	for _, cancion := range catalogoCanciones {
		// Si hay duplicados, mantener el primero
		if _, existe := mapaCatalogo[cancion.Titulo]; !existe {
			mapaCatalogo[cancion.Titulo] = cancion
		}
	}

	contadorGeneros := make(map[string]int)
	contadorArtistas := make(map[string]int)
	contadorIdiomas := make(map[string]int)

	for _, reproduccion := range reproduccionesUsuario {
		// Buscar los metadatos de la canción reproducida en el mapa del catálogo
		if cancion, existe := mapaCatalogo[reproduccion.Titulo]; existe {
			// Incrementar los contadores correspondientes
			contadorGeneros[cancion.Genero]++
			contadorArtistas[cancion.Artista]++
			contadorIdiomas[cancion.Idioma]++
		}
	}

	prefsGeneros := make([]modelos.PreferenciaGenero, 0, len(contadorGeneros))
	for genero, count := range contadorGeneros {
		prefsGeneros = append(prefsGeneros, modelos.PreferenciaGenero{
			NombreGenero:       genero,
			NumeroPreferencias: count,
		})
	}

	prefsArtistas := make([]modelos.PreferenciaArtista, 0, len(contadorArtistas))
	for artista, count := range contadorArtistas {
		prefsArtistas = append(prefsArtistas, modelos.PreferenciaArtista{
			NombreArtista:      artista,
			NumeroPreferencias: count,
		})
	}

	prefsIdiomas := make([]modelos.PreferenciaIdioma, 0, len(contadorIdiomas))
	for idioma, count := range contadorIdiomas {
		prefsIdiomas = append(prefsIdiomas, modelos.PreferenciaIdioma{
			NombreIdioma:       idioma,
			NumeroPreferencias: count,
		})
	}

	sort.Slice(prefsGeneros, func(i, j int) bool {
		return prefsGeneros[i].NumeroPreferencias > prefsGeneros[j].NumeroPreferencias
	})

	sort.Slice(prefsArtistas, func(i, j int) bool {
		return prefsArtistas[i].NumeroPreferencias > prefsArtistas[j].NumeroPreferencias
	})

	sort.Slice(prefsIdiomas, func(i, j int) bool {
		return prefsIdiomas[i].NumeroPreferencias > prefsIdiomas[j].NumeroPreferencias
	})

	respuesta := modelos.PreferenciasRespuesta{
		IdUsuario:            idUsuario,
		PreferenciasGeneros:  prefsGeneros,
		PreferenciasArtistas: prefsArtistas,
		PreferenciasIdiomas:  prefsIdiomas,
	}

	fmt.Println("CalculadorPreferencias: Cálculo finalizado")
	return respuesta
}
