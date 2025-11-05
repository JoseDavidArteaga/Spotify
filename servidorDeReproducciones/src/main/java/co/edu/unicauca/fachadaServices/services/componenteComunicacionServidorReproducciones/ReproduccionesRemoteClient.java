package co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOInput;
import feign.Headers;
import feign.Param;
import feign.RequestLine;

import java.util.List;

public interface ReproduccionesRemoteClient {

    @RequestLine("GET /reproducciones?idUsuario={idUsuario}")
    @Headers("Accept: application/json")

    List<ReproduccionesDTOInput> obtenerReproducciones(@Param("idUsuario")Integer idUsuario);

}


