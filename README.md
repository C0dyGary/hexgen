# Proyecto Hexgen

Este proyecto es una herramienta de generación de código en Go, diseñada para facilitar la creación de entidades, controladores, servicios y otros componentes siguiendo una arquitectura limpia y modular.

## Tabla de Contenidos

- [Estructura del Proyecto](#estructura-del-proyecto)
- [Descripción de Archivos](#descripción-de-archivos)
- [Instalación](#instalación)
- [Uso](#uso)
- [Contribuciones](#contribuciones)
- [Licencia](#licencia)

## Estructura del Proyecto

```plaintext
hexgen/
├── go.mod
├── go.sum
├── main.go
├── README.md
├── prueba/
├── src/
│   ├── cmd/
│   │   ├── cli.go
│   │   ├── handlers/
│   │   │   ├── genEntity.go
│   │   │   ├── genMkdir.go
│   │   │   └── genTemp.go
│   │   ├── template/
│   │       ├── load.go
│   │       ├── tmpl.go
│   │       └── templates/
│   │           ├── entity.tmpl
│   │           ├── handler.tmpl
│   │           ├── handlerMeth.tmpl
│   │           ├── port.tmpl
│   │           ├── repository.tmpl
│   │           ├── routes.tmpl
│   │           └── service.tmpl
│   ├── pkg/
│       ├── domain/
│           ├── Entity.go
│           ├── field.go
│           ├── handler.go
│           └── method.go
```

## Descripción de Archivos

### Archivos Principales

- **`main.go`**: Punto de entrada principal del programa.
- **`go.mod` y `go.sum`**: Archivos de configuración y dependencias del proyecto Go.

### Directorio `src/cmd`

- **`cli.go`**: Define la interfaz de línea de comandos para interactuar con la herramienta.

#### Subdirectorio `handlers`

- **`genEntity.go`**: Genera entidades basadas en las especificaciones proporcionadas.
- **`genMkdir.go`**: Crea directorios necesarios para la estructura del proyecto.
- **`genTemp.go`**: Maneja la generación de archivos a partir de plantillas.

#### Subdirectorio `template`

- **`load.go`**: Carga y gestiona las plantillas desde el sistema de archivos.
- **`tmpl.go`**: Proporciona funciones para procesar y renderizar plantillas.

##### Subdirectorio `templates`

- **`entity.tmpl`**: Plantilla para generar entidades.
- **`handler.tmpl`**: Plantilla para generar controladores.
- **`handlerMeth.tmpl`**: Plantilla para métodos de controladores.
- **`port.tmpl`**: Plantilla para definir puertos.
- **`repository.tmpl`**: Plantilla para repositorios.
- **`routes.tmpl`**: Plantilla para rutas.
- **`service.tmpl`**: Plantilla para servicios.

### Directorio `src/pkg/domain`

- **`Entity.go`**: Define la estructura base de una entidad.
- **`field.go`**: Maneja los campos de las entidades.
- **`handler.go`**: Define los controladores base.
- **`method.go`**: Proporciona métodos para las entidades y controladores.

## Instalación

1. Clona este repositorio:

   ```bash
   git clone https://github.com/<usuario>/<nombre-del-repo>.git
   cd <nombre-del-repo>
   ```

2. Instala las dependencias:

   ```bash
   go mod tidy
   ```

3. Ejecuta la aplicación:

   ```bash
   go run main.go
   ```

## Uso

Utiliza la herramienta para generar componentes de tu proyecto ejecutando comandos desde la línea de comandos. Por ejemplo:

```bash
go run main.go generate entity --name=User
```

## Contribuciones

¡Las contribuciones son bienvenidas! Por favor, abre un issue o envía un pull request con tus mejoras.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.
