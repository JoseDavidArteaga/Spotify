package co.edu.unicauca.infoii.correo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class CorreoApplication {

	public static void main(String[] args) {
		SpringApplication.run(CorreoApplication.class, args);
		System.out.println("Servidor de correo iniciado correctamente");
		System.out.println("¡Bienvenido al sistema de envío de correos!");
		System.out.println("Esperando peticiones de notificación...");
		System.out.println("Para salir, presione Ctrl+C");
		
	}

}
