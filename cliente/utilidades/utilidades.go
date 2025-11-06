package utilidades

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	pb "servidor.local/grpc-servidor/serviciosStreaming" //hay un error
)

func LlamarPreferencias(userID int) {
	url := fmt.Sprintf("http://localhost:2021/preferencias/calcular?idUsuario=%d", userID)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error llamando al servidor de preferencias: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("\nPreferencias recibidas:")
	fmt.Println(string(body))
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
