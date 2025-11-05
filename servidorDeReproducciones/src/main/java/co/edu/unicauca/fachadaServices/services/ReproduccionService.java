package co.edu.unicauca.fachadaServices.services;

import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOInput;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class ReproduccionService {

    private List<ReproduccionesDTOInput> lista = new ArrayList<>();

    public void registrar(ReproduccionesDTOInput dto) {
        lista.add(dto);
        System.out.println("[ECO][Java] Registro guardado. Total reproducciones: " + lista.size());
    }

    public List<ReproduccionesDTOInput> listarPorUsuario(String userId) {
        List<ReproduccionesDTOInput> res = new ArrayList<>();
        for (ReproduccionesDTOInput r : lista) {
            if (r.getUserId().equals(userId)) {
                res.add(r);
            }
        }
        return res;
    }
}

