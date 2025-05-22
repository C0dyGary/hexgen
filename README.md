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

El binario `hexgen` es una herramienta de línea de comandos (CLI) diseñada para automatizar la generación de código en proyectos Go. Su propósito principal es crear entidades, controladores, servicios, repositorios y otros componentes siguiendo una arquitectura modular y limpia.

### Ejecución del Binario

El binario se ejecuta utilizando el siguiente formato básico:
```bash
./hexgen <comando> [opciones]
```

### Comandos Disponibles

- `generate entity`: Genera una nueva entidad con sus atributos.
- `generate handler`: Crea un controlador asociado a una entidad.
- `generate service`: Genera un servicio para manejar la lógica de negocio.
- `generate repository`: Crea un repositorio para interactuar con la base de datos.
- `generate routes`: Configura las rutas HTTP para una entidad.

### Opciones de los Comandos

Cada comando puede aceptar opciones específicas. Por ejemplo:
```bash
./hexgen generate entity --name=User --fields="id:int,name:string,email:string"
```
En este caso:
- `--name`: Especifica el nombre de la entidad.
- `--fields`: Define los campos de la entidad, con su nombre y tipo.

### Estructura Generada

Al ejecutar un comando, el binario genera los archivos correspondientes en las ubicaciones adecuadas dentro del proyecto. Por ejemplo:
- Una entidad se genera en `src/pkg/domain/`.
- Un controlador se genera en `src/cmd/handlers/`.
- Las rutas se configuran en `src/cmd/template/templates/routes.tmpl`.

### Plantillas Personalizables

El binario utiliza plantillas ubicadas en `src/cmd/template/templates/` para generar el código. Estas plantillas pueden ser personalizadas para adaptarse a las necesidades específicas del proyecto.

### Ejemplo de Uso

1. Generar una nueva entidad llamada `Product` con campos específicos:
   ```bash
   ./hexgen generate entity --name=Product --fields="id:int,name:string,price:float"
   ```

2. Crear un controlador para la entidad `Product`:
   ```bash
   ./hexgen generate handler --name=Product
   ```

3. Configurar las rutas para la entidad `Product`:
   ```bash
   ./hexgen generate routes --name=Product
   ```

### Notas

- Asegúrate de que el binario `hexgen` tenga permisos de ejecución. Si no los tiene, puedes otorgarlos con:
  ```bash
  chmod +x hexgen
  ```
- Los archivos generados deben ser revisados y ajustados según las necesidades específicas del proyecto.
- El binario debe estar ubicado en un directorio incluido en tu `$PATH` para ejecutarlo desde cualquier ubicación. Si no es así, puedes moverlo a `/usr/local/bin`:
  ```bash
  sudo mv hexgen /usr/local/bin
  ```

## Contribuciones

¡Las contribuciones son bienvenidas! Por favor, abre un issue o envía un pull request con tus mejoras.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.
