package vistas

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	pbStream "servidor.local/grpc-servidor/serviciosStreaming"    //tengo error aqui
	pbSong "servidor.local/grpc-servidorCancion/serviciosCancion" //tengo error aqui

	util "cliente.local/grpc-cliente/utilidades"
)

var reader = bufio.NewReader(os.Stdin)

// / MostrarMenuPrincipal is the main entry point for the client's UI loop.
func MostrarMenuPrincipal(
	clienteCanciones pbSong.ServiciosCancionesClient,
	clienteStreaming pbStream.AudioServiceClient,
	ctx context.Context,
	nickname string,
	idUsuario int,
) {
	fmt.Printf("\n¡Bienvenido, %s!\n", nickname)

	for {
		opcion := mostrarMenuPrincipalYObtenerOpcion()

		switch opcion {
		case 1:
			// Explore musical genres
			explorarGeneros(clienteCanciones, clienteStreaming, ctx)
		case 2:
			// View user preferences
			util.LlamarPreferencias(idUsuario)
		case 3:
			// Exit
			fmt.Println("\n👋 ¡Gracias por usar Spotify Pirata!")
			return
		}
	}
}

// mostrarMenuPrincipalYObtenerOpcion shows the main menu and gets a valid choice.
func mostrarMenuPrincipalYObtenerOpcion() int {
	const minOpcion = 1
	const maxOpcion = 3 // CRITICAL FIX: Was '2', which made option 3 unreachable.

	for {
		fmt.Println("\n" + strings.Repeat("*", 50))
		fmt.Println("SPOTIFY PIRATA - MENÚ PRINCIPAL")
		fmt.Println(strings.Repeat("*", 50))
		fmt.Println("1. Escoge una canción a reproducir")
		fmt.Println("2. Ver recomendaciones de preferencias")
		fmt.Println("3. Salir")
		fmt.Print("\n📝 Seleccione una opción (1-3): ")

		input, err := leerEntradaSinEspacios()
		if err != nil {
			fmt.Println("❌ Error leyendo entrada. Intente nuevamente.")
			continue
		}

		opcion, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Por favor, ingrese un número válido.")
			continue
		}

		if opcion >= minOpcion && opcion <= maxOpcion {
			return opcion // Valid option (1, 2, or 3)
		}

		fmt.Printf("❌ Opción fuera de rango. Seleccione de %d a %d.\n", minOpcion, maxOpcion)
	}
}

// leerEntradaSinEspacios lee una línea de entrada y la limpia de espacios y saltos de línea.
func leerEntradaSinEspacios() (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// explorarGeneros maneja la exploración de géneros musicales.
func explorarGeneros(
	clienteCanciones pbSong.ServiciosCancionesClient,
	clienteStreaming pbStream.AudioServiceClient,
	ctx context.Context,
) {
	fmt.Println("\n📡 Obteniendo lista de géneros disponibles...")

	// añade timeout para la llamada gRPC
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	respuestaGeneros, err := clienteCanciones.ListarGeneros(ctxTimeout, &pbSong.Vacio{})
	if err != nil {
		fmt.Printf("❌ Error obteniendo géneros: %v\n", err)
		presionarEnterParaContinuar()
		return
	}

	if len(respuestaGeneros.Generos) == 0 {
		fmt.Println("😔 No hay géneros disponibles en este momento.")
		presionarEnterParaContinuar()
		return
	}

	// Loop para selección de géneros
	for {
		idGenero := mostrarGenerosYSeleccionar(respuestaGeneros.Generos)
		if idGenero == -1 {
			return
		}

		genero := buscarGeneroPorId(respuestaGeneros.Generos, idGenero)
		if genero == nil {
			fmt.Printf("❌ Género con ID %d no encontrado. Intente nuevamente.\n", idGenero)
			continue // Ask for genre ID again
		}

		// If a valid genre is found, enter the song exploration menu
		explorarCancionesPorGenero(clienteCanciones, clienteStreaming, ctx, genero)
	}
}

// mostrarGenerosYSeleccionar displays the list of genres and asks for a selection.
// Returns a valid genre ID or -1 to go back.
func mostrarGenerosYSeleccionar(generos []*pbSong.Genero) int32 {
	for {
		fmt.Println("\n" + strings.Repeat("*", 40))
		fmt.Println("GÉNEROS MUSICALES DISPONIBLES")
		fmt.Println(strings.Repeat("*", 40))

		for _, g := range generos {
			fmt.Printf("🎵 %d. %s\n", g.Id, g.Nombre)
		}
		fmt.Printf("0. Volver al menú principal\n")
		fmt.Print("\nSeleccione un género (por ID): ")

		input, err := leerEntradaSinEspacios()
		if err != nil {
			fmt.Println("❌ Error leyendo entrada. Intente nuevamente.")
			continue
		}

		if input == "0" {
			return -1
		}

		idGenero, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Por favor, ingrese un número válido.")
			continue
		}

		return int32(idGenero)
	}
}

// buscarGeneroPorId es una función auxiliar para encontrar un género por su ID.
func buscarGeneroPorId(generos []*pbSong.Genero, id int32) *pbSong.Genero {
	for _, g := range generos {
		if g.Id == id {
			return g
		}
	}
	return nil // Not found
}

func explorarCancionesPorGenero(
	clienteCanciones pbSong.ServiciosCancionesClient,
	clienteStreaming pbStream.AudioServiceClient,
	ctx context.Context,
	genero *pbSong.Genero,
) {
	fmt.Printf("\nBuscando canciones del género '%s'...\n", genero.Nombre)

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	respuestaCanciones, err := clienteCanciones.ListarCancionesPorGenero(ctxTimeout, &pbSong.IdGenero{Id: genero.Id})
	if err != nil {
		fmt.Printf("❌ Error obteniendo canciones: %v\n", err)
		presionarEnterParaContinuar()
		return
	}

	if len(respuestaCanciones.Canciones) == 0 {
		fmt.Printf("No se encontraron canciones para el género '%s'.\n", genero.Nombre)
		presionarEnterParaContinuar()
		return
	}

	for {
		mostrarCancionesDelGenero(respuestaCanciones.Canciones, genero.Nombre)

		cancionSeleccionada := solicitarCancionPorTituloOID(respuestaCanciones.Canciones)
		if cancionSeleccionada == nil {
			return // User chose to go back
		}

		mostrarDetallesCancion(cancionSeleccionada)

		if confirmarReproduccion() {
			reproducirCancion(clienteStreaming, ctx, cancionSeleccionada)
		}
	}
}

func mostrarCancionesDelGenero(canciones []*pbSong.Cancion, nombreGenero string) {
	fmt.Println("\n" + strings.Repeat("*", 50))
	fmt.Printf("CANCIONES DEL GÉNERO: %s\n", strings.ToUpper(nombreGenero))
	fmt.Println(strings.Repeat("*", 50))

	for i, c := range canciones {
		fmt.Printf("🎶 %d. %s - %s\n", i+1, c.Titulo, c.Artista)
	}
	fmt.Println("\nPara reproducir, escriba el número (ej: '1') o el título exacto.")
}

func solicitarCancionPorTituloOID(canciones []*pbSong.Cancion) *pbSong.Cancion {
	for {
		fmt.Print("\n📝 Ingrese el número o título de la canción (o 'volver' para regresar): ")

		input, err := leerEntradaSinEspacios()
		if err != nil {
			fmt.Println("❌ Error leyendo entrada. Intente nuevamente.")
			continue
		}

		if strings.ToLower(input) == "volver" {
			return nil // Signal to go back
		}

		if num, err := strconv.Atoi(input); err == nil {
			if num >= 1 && num <= len(canciones) {
				return canciones[num-1] // Found by index
			}
			fmt.Println("❌ Número fuera de rango. Intente nuevamente.")
			continue
		}

		// If not a number, try to match by title (case-insensitive)
		for _, c := range canciones {
			if strings.EqualFold(c.Titulo, input) {
				return c // Found by title
			}
		}

		fmt.Println("❌ No se encontró ninguna canción con ese número o título. Intente nuevamente.")
	}
}

func mostrarDetallesCancion(cancion *pbSong.Cancion) {
	fmt.Println("\n" + strings.Repeat("=", 45))
	fmt.Println("🎵 DETALLES DE LA CANCIÓN")
	fmt.Println(strings.Repeat("=", 45))
	fmt.Printf("🎶 Título: %s\n", cancion.Titulo)
	fmt.Printf("🎤 Artista: %s\n", cancion.Artista)
	fmt.Printf("📅 Año: %d\n", cancion.AnioLanzamiento)
	fmt.Printf("⏱️  Duración: %s\n", cancion.Duracion)
	// Check if ObjGenero is nil to prevent panic
	if cancion.ObjGenero != nil {
		fmt.Printf("🎸 Género: %s\n", cancion.ObjGenero.Nombre)
	}
	fmt.Printf("🗣️  Idioma: %s\n", cancion.Idioma)
	fmt.Println(strings.Repeat("=", 45))
}

// confirmarReproduccion asks the user for a yes/no confirmation.
func confirmarReproduccion() bool {
	for {
		fmt.Print("\n¿Desea reproducir esta canción? (s/n): ")

		input, err := leerEntradaSinEspacios()
		if err != nil {
			fmt.Println("❌ Error leyendo entrada. Intente nuevamente.")
			continue
		}

		input = strings.ToLower(input)

		switch input {
		case "s", "si", "sí", "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("❌ Por favor, responda 's' para sí o 'n' para no.")
		}
	}
}

// reproducirCancion handles the gRPC streaming and playback.
// It uses goroutines and channels to manage audio reception, playback, and user interruption.
func reproducirCancion(clienteStreaming pbStream.AudioServiceClient, ctx context.Context, cancion *pbSong.Cancion) {
	fmt.Printf("\nIniciando reproducción de '%s'...\n", cancion.Titulo)

	// Create a cancellable context for this specific stream
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel() // Ensures cancellation signal is sent when function exits

	stream, err := clienteStreaming.EnviarCancionMedianteStream(ctxCancel, &pbStream.PeticionDTO{
		Id:      cancion.Id,
		Formato: "mp3", // Assuming MP3 format
	})
	if err != nil {
		fmt.Printf("❌ Error iniciando streaming: %v\n", err)
		presionarEnterParaContinuar()
		return
	}

	fmt.Println("Reproduciendo canción en vivo...")
	fmt.Println("Escriba 1 para detener la reproducción.")

	// Use io.Pipe to connect the gRPC receiver to the audio player
	audioReader, audioWriter := io.Pipe()

	// Channels for synchronization
	donePlaying := make(chan struct{}) // Signals when audio playback finishes
	interruption := make(chan bool, 1) // Signals when user presses '1'
	streamError := make(chan error, 1) // Signals if the stream fails

	// Goroutine 1: Play audio from the pipe
	// This goroutine will block until audioReader receives data or is closed
	go func() {
		util.ReproducirCancion(audioReader, donePlaying)
	}()

	// Goroutine 2: Listen for keyboard input to stop
	// This goroutine runs in a loop, using the GLOBAL reader
	go func() {
		for {
			// Read from the global reader, which is free because the main loop is blocked
			input, err := reader.ReadString('\n')
			if err != nil {
				// If the main context is cancelled, this might error.
				return
			}
			if strings.TrimSpace(input) == "1" {
				interruption <- true
				return // Stop this goroutine
			}
			// If input is not '1', the loop continues, ready to read again.
		}
	}()

	go func() {
		err := util.RecibirCancion(stream, audioWriter)
		if err != nil {
			streamError <- err // Report streaming error
		}
		// When streaming is done, close the writer to signal EOF to the player
		audioWriter.Close()
	}()

	select {
	case <-interruption:
		fmt.Println("\nReproducción detenida por el usuario.")
		cancel()

	case <-donePlaying:
		fmt.Println("\nReproducción finalizada.")

	case err := <-streamError:
		fmt.Printf("\nError durante el streaming: %v\n", err)
		cancel()
	}

	audioReader.Close()
	audioWriter.Close()

	presionarEnterParaContinuar()
}

// presionarEnterParaContinuar pausa la ejecución hasta que el usuario presione Enter.
func presionarEnterParaContinuar() {
	fmt.Print("\nPresione Enter para continuar...")
	reader.ReadString('\n')
}
