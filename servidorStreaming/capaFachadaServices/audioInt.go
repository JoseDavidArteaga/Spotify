package capafachadaservices

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func StreamAudioFile(idCancion int32, funcionParaEnviarFragmento func([]byte) error) error {
	log.Printf("Canción solicitada: %d", idCancion)

	// ✅ Obtener el directorio de trabajo actual
	wd, _ := os.Getwd()

	// ✅ Ir un nivel hacia arriba desde 'main' para llegar a la carpeta raíz del servidor
	basePath := filepath.Join(wd, "..", "canciones")

	// ✅ Construir la ruta completa del archivo
	filePath := filepath.Join(basePath, strconv.FormatInt(int64(idCancion), 10)+".mp3")

	log.Printf("Intentando abrir archivo: %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 64*1024) // 64 KB por fragmento
	fragmento := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			log.Println("Canción enviada completamente desde la fachada")
			break
		}

		if err != nil {
			return fmt.Errorf("error leyendo el archivo: %w", err)
		}

		fragmento++
		log.Printf("Fragmento #%d leído (%d bytes) y enviando", fragmento, n)

		// Ejecutamos la función para enviar el fragmento al cliente
		err = funcionParaEnviarFragmento(buffer[:n])
		if err != nil {
			return fmt.Errorf("error enviando fragmento #%d: %w", fragmento, err)
		}
	}
	return nil
}
