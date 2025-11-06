package modelos

// CancionDTO representa los metadatos de una canción del catálogo
type CancionDTO struct {
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Genero  string `json:"genero"`
	Idioma  string `json:"idioma"`
}

// ReproduccionDTO representa una reproducción del historial de un usuario
type ReproduccionDTO struct {
	IdUsuario int    `json:"idUsuario"`
	Titulo    string `json:"titulo"`
	FechaHora string `json:"fechaHora"`
}

// PreferenciaGenero representa la preferencia de un género musical
type PreferenciaGenero struct {
	NombreGenero       string `json:"nombreGenero"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

// PreferenciaArtista representa la preferencia de un artista
type PreferenciaArtista struct {
	NombreArtista      string `json:"nombreArtista"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

// PreferenciaIdioma representa la preferencia de un idioma
type PreferenciaIdioma struct {
	NombreIdioma       string `json:"nombreIdioma"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

// PreferenciasRespuesta es la respuesta completa con todas las preferencias calculadas
type PreferenciasRespuesta struct {
	IdUsuario            int                  `json:"idUsuario"`
	PreferenciasGeneros  []PreferenciaGenero  `json:"preferenciasGeneros"`
	PreferenciasArtistas []PreferenciaArtista `json:"preferenciasArtistas"`
	PreferenciasIdiomas  []PreferenciaIdioma  `json:"preferenciasIdiomas"`
}

// SolicitudPreferencias representa la petición para calcular preferencias
type SolicitudPreferencias struct {
	IdUsuario int `json:"idUsuario"`
}
