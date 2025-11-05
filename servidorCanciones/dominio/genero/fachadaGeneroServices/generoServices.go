package fachadaGeneroServices

import (
	"log"

	generoDTO "servidor.local/grpc-servidorCancion/dominio/genero/dto"
	"servidor.local/grpc-servidorCancion/dominio/genero/repositorio"
	pb "servidor.local/grpc-servidorCancion/serviciosCancion"
)

type GeneroServices struct{}

// ListarGeneros obtiene todos los géneros y devuelve la respuesta lista para gRPC
func ListarGeneros() *pb.RespuestaGenerosDTO {
	log.Printf("Fachada: Listando todos los géneros")
	generos := repositorio.BuscarTodosLosGeneros()

	//* Verificar si se encontraron géneros OJOOOOOOOOOOOOO
	if generos == nil {
		return &pb.RespuestaGenerosDTO{
			Mensaje:    "No se encontraron géneros",
			Codigo:     404,
			ObjGeneros: nil,
		}
	}

	var pbGeneros []*pb.Genero
	for _, g := range generos {
		pbGeneros = append(pbGeneros, generoDTO.ToPbGenero(g))
	}

	return &pb.RespuestaGenerosDTO{
		Mensaje:    "Generos listados exitosamente",
		Codigo:     200,
		ObjGeneros: pbGeneros,
	}
}

// BuscarGenero busca un genero por ID
func BuscarGenero(id int32) *pb.RespuestaGeneroDTO {
	log.Printf("Fachada: Buscando genero con ID=%d", id)
	respuesta := repositorio.BuscarGenero(id)

	if respuesta.Codigo == 200 {
		return &pb.RespuestaGeneroDTO{
			Mensaje: "Genero encontrado",
			Codigo:  200,
			Genero:  generoDTO.ToPbGenero(respuesta.ObjGenero),
		}
	}

	return &pb.RespuestaGeneroDTO{
		Mensaje: "Genero no encontrado",
		Codigo:  404,
	}
}
