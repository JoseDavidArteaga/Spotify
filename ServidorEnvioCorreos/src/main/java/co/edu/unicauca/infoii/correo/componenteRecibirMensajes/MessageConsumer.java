package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Service;
import co.edu.unicauca.infoii.correo.commons.Simulacion;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Random;

// Servicio que consume mensajes de la cola de RabbitMQ
@Service
public class MessageConsumer {

    private static final String[] frases = {
        "¡La música es el lenguaje universal!",
        "¡El fútbol no se juega con los pies, se juega con el corazón!",
        "Cada partido es una nueva oportunidad para brillar.",
        "¡Nunca dejes de correr tras tus sueños, como tras un balón!",
        "Ganar o perder, lo importante es dejarlo todo en la cancha.",
        "El esfuerzo de hoy es la victoria de mañana.",
        "Cuando el equipo juega unido, el triunfo está más cerca.",
        "¡La pasión no se entrena, se siente!",
        "En el fútbol, como en la vida, no hay imposibles.",
        "Los grandes jugadores nacen en los momentos difíciles.",
        "¡Juega con garra, juega con alma!"
    };

    // Generador de números aleatorios para seleccionar frases motivadoras
    private final Random random = new Random();

    @RabbitListener(queues = "cola_notificaciones")
    public void receiveMessage(CancionAlmacenarDTOInput cancionRecibida) {
        System.out.println("\n==============================================");
        System.out.println("Mensaje recibido de la cola");
        
        System.out.println("Datos escuchadps: " + cancionRecibida.toString());

        System.out.println("\n Simulando envío de correo electrónico...");
        Simulacion.simular(3000, "Enviando  notificación...");
        
        // Fecha y hora actual
        LocalDateTime ahora = LocalDateTime.now();
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
        String fecha = ahora.format(formatter);
        
        // Seleccionar una frase motivadora al azar
        String fraseMotivadora = frases[random.nextInt(frases.length)];

        // Imprimir el cuerpo del correo simulado
        System.out.println("\n📩 ----------- INICIO DEL CORREO -------------");
        System.out.println("🕒 Fecha de registro: " + fecha);
        System.out.println("\n👋 ¡Hola!");
        System.out.println("🎵 Una nueva canción ha sido registrada en el sistema.");
        System.out.println("\n📋 Detalles de la canción:");

        System.out.println("\n🎶 Título:  " + cancionRecibida.getTitulo());
        System.out.println("🎤 Artista: " + cancionRecibida.getArtista());
        System.out.println("🎧 Género:  " + cancionRecibida.getGenero());
        System.out.println("🗣️ Idioma:  " + cancionRecibida.getIdioma());

        System.out.println("\n💬 " + fraseMotivadora);
        System.out.println(" ------------------------------------------------");
        System.out.println(" --------------- ✅ FIN DEL CORREO -------------");
        System.out.println(" ================================================");

    }
}