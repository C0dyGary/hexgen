package handlers

import (
	"fmt"
	"github.com/C0dyGary/hexgen/src/cmd/template"

	"github.com/C0dyGary/hexgen/src/pkg/domain"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func GenEntity(c *cli.Context) error {
	entityName := c.String("name")
	fieldsRaw := c.StringSlice("fields")
	/*generateService := c.Bool("service")
	generateRepository := c.Bool("repository")
	generateServiceRepository := c.Bool("service-repository")*/
	var fields []domain.Field

	for _, f := range fieldsRaw {
		var name, typ string
		parts := strings.Split(f, ":")
		if len(parts) != 2 {
			fmt.Println("Error en el formato de los campos")
			return nil
		}
		name = parts[0]
		typ = parts[1]
		fields = append(fields, domain.Field{Name: name, Type: typ, JSON: strings.ToLower(name)})
	}

	// Carga la plantilla "entity.tmpl"
	tmpl, err := template.LoadTemplate("entity")
	if err != nil {
		fmt.Println("Error al cargar la plantilla entity.tmpl")
		return err
	}

	// Crea el archivo destino en la carpeta deseada
	file, err := os.Create(fmt.Sprintf("src/pkg/domain/%s.go", entityName))
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Print("Error al cerrar el archivo")
			panic(err)
		}
	}(file)

	entity := domain.Entity{Name: entityName, Fields: fields}
	err = tmpl.Execute(file, entity)
	if err != nil {
		fmt.Println("Error al ejecutar la plantilla")
		return err
	}
	fmt.Printf("Estructura %s creada con éxito.\n", entityName)

	/*if generateService {
		return GenService(entityName)
	}

	if generateRepository {
		return GenRepository(entityName)
	}

	if generateServiceRepository {
		err = GenService(entityName)
		if err != nil {
			return err
		}
		err = GenRepository(entityName)
		if err != nil {
			return err
		}
	}*/
	return nil
}
