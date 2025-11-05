package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import feign.Feign;
import feign.jackson.JacksonDecoder;

import java.util.ArrayList;
import java.util.List;

/**
 * Clase encargada de la comunicación con el servidor remoto de reproducciones
 * (microservicio o API externa) usando Feign Client.
 */
public class ComunicacionServidorReproducciones {

    // URL base del servidor de reproducciones remoto.
    private static final String BASE_URL = "http://localhost:2020";
    
    //Cliente Feign para realizar las peticiones al servidor de reproducciones.
    private final ReproduccionesRemoteClient client;

    /**
     * Constructor que inicializa el cliente Feign.
     * Configura el cliente para usar Jackson para la decodificación JSON
     * y apunta al BASE_URL.
     */
    public ComunicacionServidorReproducciones() {
        this.client = Feign.builder()
            .decoder(new JacksonDecoder()) // Usa Jackson para deserializar el JSON de la respuesta
            .target(ReproduccionesRemoteClient.class, BASE_URL); // Define la interfaz y la URL base
    }
    
    //Obtiene la lista de reproducciones de un usuario específico desde el servidor remoto.
    public List<ReproduccionesDTOEntrada> obtenerReproduccionesRemotas(Integer idUsuario) {
        try {
            List<ReproduccionesDTOEntrada> reproducciones = client.obtenerReproducciones(idUsuario);
            
            // Verifica si la lista es nula y devuelve una lista vacía en su lugar.
            if (reproducciones == null) {
                return new ArrayList<>();
            }
            
            return reproducciones; 
        } catch (Exception e) {
            // En caso de cualquier error de comunicación, se maneja la excepción 
            // y se retorna una lista vacía.
            return new ArrayList<>();
        }
    }
    
}


