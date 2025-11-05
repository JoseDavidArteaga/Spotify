package co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias;

import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

import co.edu.unicauca.fachadaServices.DTO.CancionDTOEntrada;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaArtistaDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciaGeneroDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;

/**
 * Clase que se encarga de calcular las preferencias de géneros y artistas
 * a partir de una lista de canciones y reproducciones.
 */
public class CalculadorPreferencias {

    /**
     * Calcula las preferencias de géneros y artistas para un usuario específico.
     *
     * @param idUsuario El ID del usuario para el que se calculan las preferencias.
     * @param canciones La lista de todas las canciones disponibles.
     * @param reproducciones La lista de reproducciones registradas.
     * @return Un objeto PreferenciasDTORespuesta con las preferencias calculadas.
     */
    public PreferenciasDTORespuesta calcular(Integer idUsuario,
                                             List<CancionDTOEntrada> canciones,
                                             List<ReproduccionesDTOEntrada> reproducciones) {
        // 1. Crear un mapa de canciones para búsqueda rápida por ID
        Map<Integer, CancionDTOEntrada> mapaCanciones = canciones.stream()
            .filter(Objects::nonNull) // Filtrar nulos
            .filter(c -> c.getId() != null) // Filtrar canciones sin ID
            // Recolectar en un mapa, usando el ID como clave. Si hay duplicados, se queda con el primero (a).
            .collect(Collectors.toMap(CancionDTOEntrada::getId, c -> c, (a, b) -> a));

        // 2. Inicializar contadores para géneros y artistas
        Map<String, Integer> contadorGeneros = new HashMap<>();
        Map<String, Integer> contadorArtistas = new HashMap<>();

        // 3. Iterar sobre las reproducciones para contar preferencias
        for (ReproduccionesDTOEntrada r : reproducciones) {
            Integer idCancion = r.getIdCancion();
            if (idCancion == null) continue; // Saltar si no hay ID de canción

            CancionDTOEntrada c = mapaCanciones.get(idCancion);
            if (c == null) {
                continue; // Saltar si la canción no se encuentra en el mapa
            }

            // Obtener género y artista, usando "Desconocido" si son nulos
            String genero = c.getGenero() == null ? "Desconocido" : c.getGenero();
            String artista = c.getArtista() == null ? "Desconocido" : c.getArtista();

            // Incrementar contadores
            contadorGeneros.put(genero, contadorGeneros.getOrDefault(genero, 0) + 1);
            contadorArtistas.put(artista, contadorArtistas.getOrDefault(artista, 0) + 1);
        }

        // 4. Transformar y ordenar las preferencias de géneros
        List<PreferenciaGeneroDTORespuesta> preferenciasGeneros = contadorGeneros.entrySet().stream()
            .map(e -> {
                PreferenciaGeneroDTORespuesta dto = new PreferenciaGeneroDTORespuesta();
                dto.setNombreGenero(e.getKey());
                dto.setNumeroPreferencias(e.getValue());
                return dto;
            })
            // Ordenar por número de preferencias descendente, y por nombre de género ascendente como desempate
            .sorted(Comparator.comparingInt(PreferenciaGeneroDTORespuesta::getNumeroPreferencias).reversed()
                            .thenComparing(PreferenciaGeneroDTORespuesta::getNombreGenero))
            .collect(Collectors.toList());

        // 5. Transformar y ordenar las preferencias de artistas
        List<PreferenciaArtistaDTORespuesta> preferenciasArtistas = contadorArtistas.entrySet().stream()
            .map(e -> {
                PreferenciaArtistaDTORespuesta dto = new PreferenciaArtistaDTORespuesta();
                dto.setNombreArtista(e.getKey());
                dto.setNumeroPreferencias(e.getValue());
                return dto;
            })
            // Ordenar por número de preferencias descendente, y por nombre de artista ascendente como desempate
            .sorted(Comparator.comparingInt(PreferenciaArtistaDTORespuesta::getNumeroPreferencias).reversed()
                            .thenComparing(PreferenciaArtistaDTORespuesta::getNombreArtista))
            .collect(Collectors.toList());

        // 6. Construir el objeto de respuesta final
        PreferenciasDTORespuesta respuesta = new PreferenciasDTORespuesta();
        respuesta.setIdUsuario(idUsuario);
        respuesta.setPreferenciasGeneros(preferenciasGeneros);
        respuesta.setPreferenciasArtistas(preferenciasArtistas);

        return respuesta;
    }
}