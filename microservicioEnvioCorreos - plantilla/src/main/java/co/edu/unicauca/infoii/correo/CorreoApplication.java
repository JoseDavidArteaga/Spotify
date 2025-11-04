package co.edu.unicauca.infoii.correo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class CorreoApplication {

	public static void main(String[] args) {
		SpringApplication.run(CorreoApplication.class, args);
		System.out.println("Microservicio de envío de correos iniciado correctamente.");
		
	}

}
