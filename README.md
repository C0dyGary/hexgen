Parece que deseas un `README` que abarque todos los archivos y su propósito en el proyecto. Aquí tienes un ejemplo actualizado que incluye una descripción general de los archivos mencionados:

```markdown
# Proyecto Go: API REST con Fiber

Este proyecto es una API REST desarrollada en Go utilizando el framework Fiber. La arquitectura sigue un enfoque modular y limpio, con capas bien definidas para manejar rutas, controladores, servicios y repositorios.

## Tabla de Contenidos

- [Requisitos](#requisitos)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Descripción de Archivos](#descripción-de-archivos)
- [Instalación](#instalación)
- [Uso](#uso)
- [Contribuciones](#contribuciones)
- [Licencia](#licencia)

## Requisitos

- Go 1.20 o superior
- Base de datos compatible con GORM (MySQL, PostgreSQL, SQLite, etc.)
- [Go Modules](https://blog.golang.org/using-go-modules) habilitado

## Estructura del Proyecto

```plaintext
src/
├── cmd/
│   ├── api/
│   │   ├── handler/         # Controladores para manejar las solicitudes HTTP
│   │   └── routes/          # Definición de rutas
│   ├── template/            # Plantillas para generación de código
├── pkg/
│   ├── repository/          # Repositorios para interactuar con la base de datos
│   ├── service/             # Lógica de negocio
├── go.mod                   # Archivo de dependencias
```

## Descripción de Archivos

### `src/cmd/template/load.go`
Este archivo carga plantillas desde el sistema de archivos embebido utilizando `embed`. Permite generar código dinámico basado en plantillas `.tmpl`.

### `src/cmd/template/templates/routes.tmpl`
Plantilla para definir rutas de una entidad. Genera código que configura rutas HTTP para operaciones CRUD.

### `src/cmd/template/templates/handler.tmpl`
Plantilla para generar controladores (`Handler`) de una entidad. Define la estructura del controlador y su dependencia en el servicio correspondiente.

### `src/cmd/api/handler/`
Contiene los controladores generados para manejar las solicitudes HTTP. Cada controlador está asociado a una entidad específica.

### `src/pkg/repository/`
Define los repositorios que interactúan directamente con la base de datos. Implementa las operaciones CRUD.

### `src/pkg/service/`
Contiene la lógica de negocio. Los servicios actúan como intermediarios entre los controladores y los repositorios.

### `go.mod`
Archivo de configuración de dependencias del proyecto. Define los módulos y versiones requeridas.

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

3. Configura la base de datos en el archivo correspondiente.

4. Ejecuta la aplicación:

   ```bash
   go run main.go
   ```

## Uso

La API expone endpoints para realizar operaciones CRUD sobre las entidades definidas. Las rutas se generan dinámicamente a partir de las plantillas.

## Contribuciones

¡Las contribuciones son bienvenidas! Por favor, abre un issue o envía un pull request con tus mejoras.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.
```