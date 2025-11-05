package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorCanciones;

import co.edu.unicauca.fachadaServices.DTO.CancionDTOEntrada;
import feign.Feign;
import feign.jackson.JacksonDecoder;

import java.util.ArrayList;
import java.util.List;

import org.springframework.stereotype.Component;

/**
 * Componente Spring encargado de la comunicación con el servidor remoto de canciones
 * (microservicio o API externa) usando Feign Client.
 */
@Component
public class ComunicacionServidorCanciones {

    //URL base del servidor de canciones remoto.
    private static final String BASE_URL = "http://localhost:5051";
    
    //Cliente Feign para realizar las peticiones al servidor de canciones.
    private final CancionesRemoteClient client;

    /**
     * Constructor que inicializa el cliente Feign.
     * Configura el cliente para usar Jackson para la decodificación JSON
     * y apunta al BASE_URL.
     */
    public ComunicacionServidorCanciones() {
        this.client = Feign.builder()
            .decoder(new JacksonDecoder()) // Usa Jackson para deserializar el JSON de la respuesta
            .target(CancionesRemoteClient.class, BASE_URL); // Define la interfaz y la URL base
    }

    /**
     * Obtiene la lista de canciones desde el servidor remoto.
     *
     * @return Una lista de {@code CancionDTOEntrada}. Retorna una lista vacía
     * ({@code ArrayList}) en caso de que la respuesta sea nula o si ocurre
     * alguna excepción durante la comunicación.
     */
    public List<CancionDTOEntrada> obtenerCancionesRemotas() {
        try {
            List<CancionDTOEntrada> canciones = client.obtenerCanciones();
            if (canciones == null) {
                // Si la respuesta es nula, se devuelve una lista vacía
                return new ArrayList<>();
            }
            
            // Si la lista no es nula, se devuelve la lista obtenida
            return canciones; 
        } catch (Exception e) {
            // En caso de cualquier error de comunicación (conexión, timeout, etc.),
            // se maneja la excepción y se retorna una lista vacía para no romper el flujo.
            return new ArrayList<>();
        }
    }
}