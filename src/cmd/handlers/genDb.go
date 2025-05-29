package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/C0dyGary/hexgen/src/cmd/template"
	"github.com/C0dyGary/hexgen/src/pkg/domain"
	"github.com/urfave/cli/v2"
)

func GenDb(c *cli.Context) error {
	entityName := c.String("name")
	project := c.String("project")
	db := c.String("database")

	if err := os.MkdirAll(fmt.Sprintf("src/pkg/repository/%s/%s", db, strings.ToLower(entityName)), os.ModePerm); err != nil {
		fmt.Println("Error al crear el directorio")
		return err
	}
	if err := RepositoryDB(entityName, "*mongo.Database", db); err != nil {
		fmt.Println("Error al crear el archivo repositorio")
		return err
	}
	fmt.Printf("Repository %s created successfully.\n", entityName)

	for _, method := range []string{"insert", "select", "selectAll", "update", "delete"} {
		if err := RepositoryFunctionsDB(entityName, method, project, db); err != nil {
			fmt.Println("Error al crear el archivo")
			return err
		}
	}
	fmt.Printf("Repository functions %s created successfully.\n", entityName)
	return nil
}
func RepositoryFunctionsDB(entityName, method, project, db string) error {
	file, err := os.Create(fmt.Sprintf("src/pkg/repository/%s/%s/%s.go", db, strings.ToLower(entityName), method))
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

	if _, err = file.WriteString(fmt.Sprintf("package %s\n\nimport \"github.com/C0dyGary/%s/src/pkg/domain\"\n", strings.ToLower(entityName), project)); err != nil {
		fmt.Println("Error al escribir en el archivo 1")
		return err
	}
	if _, err = file.WriteString(template.GenerateSignature("repository", method, entityName)); err != nil {
		fmt.Println("Error al escribir en el archivo 2")
		return err
	}
	return nil
}

func RepositoryDB(entityName, typeDb, db string) error {
	file, err := os.Create(fmt.Sprintf("src/pkg/repository/%s/%s/repository.go", db, strings.ToLower(entityName)))
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

	repository := domain.EntityRepository{NameLower: strings.ToLower(entityName), TypeDB: typeDb}
	tmpl, err := template.LoadTemplate("repositorydb")
	if err != nil {
		fmt.Println("Error al cargar la plantilla repositorydb.tmpl")
	}
	err = tmpl.Execute(file, repository)

	return nil
}
