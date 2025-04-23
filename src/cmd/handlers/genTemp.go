package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/C0dyGary/hexgen/src/cmd/template"
	"github.com/C0dyGary/hexgen/src/pkg/domain"
	"github.com/urfave/cli/v2"
)

func GenTemplate(c *cli.Context) error {
	entityName := c.String("name")
	project := c.String("project")

	file, err := os.Create(fmt.Sprintf("src/pkg/port/%s.go", strings.ToLower(entityName)))
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

	tmpl, err := template.LoadTemplate("port")
	if err != nil {
		fmt.Println("Error al cargar la plantilla port.tmpl")
		return err
	}

	objEntityName := domain.EntityName{Name: entityName, Project: project}

	err = tmpl.Execute(file, objEntityName)
	fmt.Printf("Port %s created successfully.\n", entityName)
	// ================= Create Service =========================
	if err = os.MkdirAll(fmt.Sprintf("src/pkg/service/%s", strings.ToLower(entityName)), os.ModePerm); err != nil {
		fmt.Println("Error al crear el directorio")
		return err
	}
	if err = Service(entityName, project); err != nil {
		fmt.Println("Error al crear el archivo servicio")
		return err
	}
	fmt.Printf("Service %s created successfully.\n", entityName)
	for _, method := range []string{"create", "delete", "read", "list", "update"} {
		if err := ServiceFunctions(entityName, project, method); err != nil {
			fmt.Println("Error al crear el archivo")
			return err
		}
	}
	fmt.Printf("Service functions %s created successfully.\n", entityName)

	// ================= Create Repository =========================
	if err = os.MkdirAll(fmt.Sprintf("src/pkg/repository/gorm/%s", strings.ToLower(entityName)), os.ModePerm); err != nil {
		fmt.Println("Error al crear el directorio")
		return err
	}
	if err = Repository(entityName, "*gorm.Model"); err != nil {
		fmt.Println("Error al crear el archivo repositorio")
		return err
	}
	fmt.Printf("Repository %s created successfully.\n", entityName)

	for _, method := range []string{"insert", "select", "selectAll", "update", "delete"} {
		if err := RepositoryFunctions(entityName, method, project); err != nil {
			fmt.Println("Error al crear el archivo")
			return err
		}
	}
	fmt.Printf("Repository functions %s created successfully.\n", entityName)
	// ================= Create Handler =========================
	if err = os.MkdirAll(fmt.Sprintf("src/cmd/api/handler/%s", strings.ToLower(entityName)), os.ModePerm); err != nil {
		fmt.Println("Error al crear el directorio")
		return err
	}
	if err = Handler(entityName, project); err != nil {
		fmt.Println("Error al crear el archivo handler")
		return err
	}
	fmt.Printf("Handler %s created successfully.\n", entityName)
	for _, method := range []string{"get", "post", "put", "delete", "patch", "getById"} {
		if err := HandlerFunctions(entityName, method, project); err != nil {
			fmt.Println("Error al crear el archivo")
			return err
		}
	}
	fmt.Printf("Handler functions %s created successfully.\n", entityName)
	return nil
}

func Service(entityName, project string) error {
	file, err := os.Create(fmt.Sprintf("src/pkg/service/%s/service.go", strings.ToLower(entityName)))
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

	service := domain.EntityName{Name: entityName, Project: project, NameLower: strings.ToLower(entityName)}
	tmpl, err := template.LoadTemplate("service")
	if err != nil {
		fmt.Println("Error al cargar la plantilla service.tmpl")
	}
	err = tmpl.Execute(file, service)
	if err != nil {
		fmt.Println("Error al ejecutar la plantilla")
	}

	return nil
}

func ServiceFunctions(entityName, project, method string) error {
	file, err := os.Create(fmt.Sprintf("src/pkg/service/%s/%s.go", strings.ToLower(entityName), method))
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Print("Error al cerrar el archivo\n")
			panic(err)
		}
	}(file)

	if _, err = file.WriteString(fmt.Sprintf("package %s\n\nimport \"github.com/C0dyGary/%s/src/pkg/domain\"\n", strings.ToLower(entityName), project)); err != nil {
		fmt.Println("Error al escribir en el archivo 1")
		return err
	}
	if _, err = file.WriteString(template.GenerateSignature("service", method, entityName)); err != nil {
		fmt.Println("Error al escribir en el archivo 2")
		return err
	}
	return nil
}

func Repository(entityName, typeDb string) error {
	file, err := os.Create(fmt.Sprintf("src/pkg/repository/gorm/%s/repository.go", strings.ToLower(entityName)))
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
	tmpl, err := template.LoadTemplate("repository")
	if err != nil {
		fmt.Println("Error al cargar la plantilla repository.tmpl")
	}
	err = tmpl.Execute(file, repository)

	return nil
}

func RepositoryFunctions(entityName, method, project string) error {
	file, err := os.Create(fmt.Sprintf("src/pkg/repository/gorm/%s/%s.go", strings.ToLower(entityName), method))
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

func Handler(entityName, project string) error {
	file, err := os.Create(fmt.Sprintf("src/cmd/api/handler/%s/handler.go", strings.ToLower(entityName)))
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

	handler := domain.Handler{NameLower: strings.ToLower(entityName), Project: project, Name: entityName}
	tmpl, err := template.LoadTemplate("handler")
	if err != nil {
		fmt.Println("Error al cargar la plantilla handler.tmpl")
	}
	err = tmpl.Execute(file, handler)
	if err != nil {
		fmt.Println("Error al ejecutar la plantilla")
	}

	return nil
}

func HandlerFunctions(entityName, method, project string) error {
	file, err := os.Create(fmt.Sprintf("src/cmd/api/handler/%s/%s.go", strings.ToLower(entityName), method))
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

	if _, err = file.WriteString(fmt.Sprintf("package %s\n\nimport (\n\t\"github.com/C0dyGary/%s/src/pkg/domain\"\n\t\"github.com/gofiber/fiber/v2\"\n)\n", strings.ToLower(entityName), project)); err != nil {
		fmt.Println("Error al escribir en el archivo 1")
		return err
	}
	if _, err = file.WriteString(template.GenerateHandler("handler", method, entityName)); err != nil {
		fmt.Println("Error al escribir en el archivo 2")
		return err
	}
	return nil
}
