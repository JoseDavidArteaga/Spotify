package utilidades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	pb "servidor.local/grpc-servidor/serviciosStreaming" //error
)

// PreferenciasRespuesta estructura para parsear la respuesta JSON
type PreferenciasRespuesta struct {
	IdUsuario            int                  `json:"idUsuario"`
	PreferenciasGeneros  []PreferenciaGenero  `json:"preferenciasGeneros"`
	PreferenciasArtistas []PreferenciaArtista `json:"preferenciasArtistas"`
	PreferenciasIdiomas  []PreferenciaIdioma  `json:"preferenciasIdiomas"`
}

type PreferenciaGenero struct {
	NombreGenero       string `json:"nombreGenero"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

type PreferenciaArtista struct {
	NombreArtista      string `json:"nombreArtista"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

type PreferenciaIdioma struct {
	NombreIdioma       string `json:"nombreIdioma"`
	NumeroPreferencias int    `json:"numeroPreferencias"`
}

func LlamarPreferencias(userID int) {
	url := "http://localhost:2021/preferencias/calcular"

	// Crear el JSON con el ID del usuario
	jsonData := []byte(fmt.Sprintf(`{"idUsuario": %d}`, userID))

	fmt.Println("\n📊 Consultando preferencias musicales...")

	// Hacer la petición POST con el JSON en el body
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Error llamando al servidor de preferencias: %v\n", err)
		fmt.Println("💡 Asegúrese de que el Servidor de Preferencias esté ejecutándose en el puerto 2021")
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Error leyendo respuesta: %v\n", err)
		return
	}

	// Verificar el código de estado
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("❌ Error: El servidor respondió con código %d\n", resp.StatusCode)
		fmt.Println(string(body))
		return
	}

	// Parsear el JSON
	var preferencias PreferenciasRespuesta
	if err := json.Unmarshal(body, &preferencias); err != nil {
		fmt.Printf("❌ Error parseando respuesta JSON: %v\n", err)
		fmt.Println("Respuesta recibida:", string(body))
		return
	}

	// Mostrar las preferencias de forma estructurada
	mostrarPreferenciasFormateadas(preferencias)
}

func mostrarPreferenciasFormateadas(prefs PreferenciasRespuesta) {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Printf("║     PREFERENCIAS MUSICALES - USUARIO ID: %-17d║\n", prefs.IdUsuario)
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	// Mostrar géneros favoritos
	fmt.Println("\n🎸 GÉNEROS FAVORITOS:")
	fmt.Println(strings.Repeat("─", 60))
	if len(prefs.PreferenciasGeneros) == 0 {
		fmt.Println("   No hay datos de géneros disponibles")
	} else {
		for i, genero := range prefs.PreferenciasGeneros {
			fmt.Printf("   %d. %-30s %3d reproducciones\n",
				i+1, genero.NombreGenero, genero.NumeroPreferencias)
		}
	}

	// Mostrar artistas favoritos
	fmt.Println("\n🎤 ARTISTAS FAVORITOS:")
	fmt.Println(strings.Repeat("─", 60))
	if len(prefs.PreferenciasArtistas) == 0 {
		fmt.Println("   No hay datos de artistas disponibles")
	} else {
		for i, artista := range prefs.PreferenciasArtistas {
			fmt.Printf("   %d. %-30s %3d reproducciones\n",
				i+1, artista.NombreArtista, artista.NumeroPreferencias)
		}
	}

	// Mostrar idiomas favoritos
	fmt.Println("\n🗣️  IDIOMAS FAVORITOS:")
	fmt.Println(strings.Repeat("─", 60))
	if len(prefs.PreferenciasIdiomas) == 0 {
		fmt.Println("   No hay datos de idiomas disponibles")
	} else {
		for i, idioma := range prefs.PreferenciasIdiomas {
			fmt.Printf("   %d. %-30s %3d reproducciones\n",
				i+1, idioma.NombreIdioma, idioma.NumeroPreferencias)
		}
	}

	fmt.Println("\n" + strings.Repeat("═", 60))
}

func ReproducirCancion(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Fatalf("error decodificando MP3: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(canalSincronizacion)
	})))
}

func RecibirCancion(
	stream pb.AudioService_EnviarCancionMedianteStreamClient,
	writer *io.PipeWriter,
) error {
	// Aseguramos que el 'writer' se cierre al salir de esta función,
	// sin importar cómo (éxito, error, etc.).
	// Esto es crucial para que el 'reader' (reproductor) sepa cuándo parar.
	defer writer.Close()

	noFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			// Fin del stream. Salimos del bucle. defer se encargará de cerrar.
			fmt.Println("Canción recibida completa.")
			break
		}
		if err != nil {
			// Error en el stream. Salimos. defer se encargará de cerrar.
			return fmt.Errorf("error recibiendo chunk: %v", err)
		}

		noFragmento++
		fmt.Printf("\n Fragmento #%d recibido (%d bytes) reproduciendo...", noFragmento, len(fragmento.Data))

		// Escribir los datos en el pipe para que el reproductor los lea
		if _, err := writer.Write(fragmento.Data); err != nil {
			// Este error usualmente significa que el reproductor (reader) se cerró
			// (ej: por interrupción del usuario).
			return fmt.Errorf("error escribiendo en pipe: %v", err)
		}
	}

	// ELIMINADO: <-canalSincronizacion (Esto causaba el deadlock)
	fmt.Println("Recepción de stream finalizada.")
	return nil
}
