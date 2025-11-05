package repositorio

import (
	dto "servidor.local/grpc-servidorCancion/dominio/cancion/dto"
	canciones "servidor.local/grpc-servidorCancion/dominio/cancion/modelo"
	repoGenero "servidor.local/grpc-servidorCancion/dominio/genero/repositorio"
)

// VectorCanciones ahora es un slice dinámico en vez de un arreglo fijo
var VectorCanciones []canciones.Cancion

// nextID se usa para asignar IDs únicos auto-incrementales dentro del proceso
var nextID int32 = 1

// CargarCanciones inicializa el slice con datos de ejemplo
func CargarCanciones() {
	// Salsa
	var cancion1 canciones.Cancion
	cancion1.Id = nextID
	nextID++
	cancion1.Titulo = "Pedro Navaja"
	cancion1.Autor = "Rubén Blades"
	cancion1.Duracion = "7:20"
	cancion1.AnioLanzamiento = 1978
	cancion1.Genero = repoGenero.VectorGeneros[0]
	cancion1.Idioma = "Español"

	var cancion2 canciones.Cancion
	cancion2.Id = nextID
	nextID++
	cancion2.Titulo = "La Rebelión"
	cancion2.Autor = "Joe Arroyo"
	cancion2.Duracion = "4:30"
	cancion2.AnioLanzamiento = 1988
	cancion2.Genero = repoGenero.VectorGeneros[0]
	cancion2.Idioma = "Español"

	// Merengue
	var cancion3 canciones.Cancion
	cancion3.Id = nextID
	nextID++
	cancion3.Titulo = "Tu eres Ajena"
	cancion3.Autor = "Eddy Herrera"
	cancion3.Duracion = "4:20"
	cancion3.AnioLanzamiento = 1995
	cancion3.Genero = repoGenero.VectorGeneros[1]
	cancion3.Idioma = "Español"

	var cancion4 canciones.Cancion
	cancion4.Id = nextID
	nextID++
	cancion4.Titulo = "La Bilirrubina"
	cancion4.Autor = "Juan Luis Guerra"
	cancion4.Duracion = "3:50"
	cancion4.AnioLanzamiento = 1990
	cancion4.Genero = repoGenero.VectorGeneros[1]
	cancion4.Idioma = "Español"

	// Rock
	var cancion5 canciones.Cancion
	cancion5.Id = nextID
	nextID++
	cancion5.Titulo = "De Música Ligera"
	cancion5.Autor = "Soda Stereo"
	cancion5.Duracion = "3:32"
	cancion5.AnioLanzamiento = 1990
	cancion5.Genero = repoGenero.VectorGeneros[2]
	cancion5.Idioma = "Español"

	var cancion6 canciones.Cancion
	cancion6.Id = nextID
	nextID++
	cancion6.Titulo = "Lamento Boliviano"
	cancion6.Autor = "Los Enanitos Verdes"
	cancion6.Duracion = "4:05"
	cancion6.AnioLanzamiento = 1994
	cancion6.Genero = repoGenero.VectorGeneros[2]
	cancion6.Idioma = "Español"

	// Pop
	var cancion7 canciones.Cancion
	cancion7.Id = nextID
	nextID++
	cancion7.Titulo = "Bailando"
	cancion7.Autor = "Enrique Iglesias ft. Descemer Bueno & Gente de Zona"
	cancion7.Duracion = "4:03"
	cancion7.AnioLanzamiento = 2014
	cancion7.Genero = repoGenero.VectorGeneros[3]
	cancion7.Idioma = "Español"

	var cancion8 canciones.Cancion
	cancion8.Id = nextID
	nextID++
	cancion8.Titulo = "Despacito"
	cancion8.Autor = "Luis Fonsi ft. Daddy Yankee"
	cancion8.Duracion = "3:48"
	cancion8.AnioLanzamiento = 2017
	cancion8.Genero = repoGenero.VectorGeneros[3]
	cancion8.Idioma = "Español"

	// Metal
	var cancion9 canciones.Cancion
	cancion9.Id = nextID
	nextID++
	cancion9.Titulo = "Master of Puppets"
	cancion9.Autor = "Metallica"
	cancion9.Duracion = "8:36"
	cancion9.AnioLanzamiento = 1986
	cancion9.Genero = repoGenero.VectorGeneros[4]
	cancion9.Idioma = "Inglés"

	var cancion10 canciones.Cancion
	cancion10.Id = nextID
	nextID++
	cancion10.Titulo = "The Number of the Beast"
	cancion10.Autor = "Iron Maiden"
	cancion10.Duracion = "4:50"
	cancion10.AnioLanzamiento = 1982
	cancion10.Genero = repoGenero.VectorGeneros[4]
	cancion10.Idioma = "Inglés"

	VectorCanciones = append(VectorCanciones, cancion1, cancion2, cancion3, cancion4, cancion5, cancion6, cancion7, cancion8, cancion9, cancion10)
}

// BuscarCancion busca una canción por título
func BuscarCancion(titulo string) dto.RespuestaDTO {
	var respuesta dto.RespuestaDTO
	for i := 0; i < len(VectorCanciones); i++ {
		if VectorCanciones[i].Titulo == titulo {
			respuesta.ObjCancion = VectorCanciones[i]
			respuesta.Codigo = 200
			respuesta.Mensaje = "Cancion encontrada correctamente"
			return respuesta
		}
	}
	respuesta.Codigo = 404
	respuesta.Mensaje = "La cancion no se encontro"
	return respuesta
}
