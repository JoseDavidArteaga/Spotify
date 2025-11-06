package co.edu.unicauca.infoii.correo.DTOs;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
public class CancionAlmacenarDTOInput {

    private String titulo;
    private String artista;
    private String genero;
    private String idioma;
    private String mensaje;
}