# Servidor de Preferencias Musicales

Servidor REST en Go que calcula las preferencias musicales de un usuario basándose en su historial de reproducciones.

## Descripción

Este servidor actúa como un orquestador que:
1. Obtiene el catálogo completo de canciones del Servidor de Canciones
2. Obtiene el historial de reproducciones del usuario del Servidor de Reproducciones
3. Calcula estadísticas agregadas por género, artista e idioma
4. Retorna las preferencias ordenadas por número de reproducciones

## Arquitectura

```
ServidorDePreferencias/
├── main/
│   └── main.go                          # Punto de entrada del servidor
├── controladores/
│   └── controlador_preferencias.go      # Handlers HTTP
├── servicios/
│   ├── servicio_preferencias.go         # Orquestador de lógica de negocio
│   └── calculador_preferencias.go       # Algoritmo de cálculo
├── clientes/
│   ├── cliente_canciones.go             # Cliente HTTP para Servidor de Canciones
│   └── cliente_reproducciones.go        # Cliente HTTP para Servidor de Reproducciones
├── modelos/
│   └── dtos.go                          # Estructuras de datos (DTOs)
└── go.mod                               # Dependencias
```

## Dependencias

- **Servidor de Canciones**: `http://localhost:5000`
- **Servidor de Reproducciones**: `http://localhost:5002`

## Endpoints

### POST /preferencias/calcular

Calcula las preferencias musicales de un usuario.

**Request:**
```json
{
  "idUsuario": 1
}
```

**Response:**
```json
{
  "idUsuario": 1,
  "preferenciasGeneros": [
    {
      "nombreGenero": "Rock",
      "numeroPreferencias": 15
    },
    {
      "nombreGenero": "Pop",
      "numeroPreferencias": 10
    }
  ],
  "preferenciasArtistas": [
    {
      "nombreArtista": "The Beatles",
      "numeroPreferencias": 8
    }
  ],
  "preferenciasIdiomas": [
    {
      "nombreIdioma": "Inglés",
      "numeroPreferencias": 20
    }
  ]
}
```

## Instalación y Ejecución

1. Asegúrate de tener Go instalado (versión 1.16 o superior)

2. Navega al directorio del servidor:
```bash
cd ServidorDePreferencias
```

3. Descarga las dependencias (si es necesario):
```bash
go mod tidy
```

4. Ejecuta el servidor:
```bash
go run main/main.go
```

El servidor estará disponible en `http://localhost:2021`

## Prueba con curl

```bash
curl -X POST http://localhost:2021/preferencias/calcular \
  -H "Content-Type: application/json" \
  -d '{"idUsuario": 1}'
```

## Notas

- El servidor debe iniciarse **después** de los servidores de Canciones y Reproducciones
- Las preferencias se calculan en tiempo real basándose en los datos actuales
- Los resultados están ordenados de mayor a menor número de reproducciones
