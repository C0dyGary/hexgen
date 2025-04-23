package template

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type Config struct {
	importTemplate string
	receiverPrefix string
	funcPrefix     string
	params         string
	returnType     string
	returnTemplate string
}

var templates = map[string]map[string]Config{
	"service": {
		"create": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Create",
			params:         "(%s domain.%s)",
			returnType:     "(*domain.%s, error)",
			returnTemplate: "return s.Repository.Insert{{.EntityName}}({{.VariableName}})",
		},
		"read": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Read",
			params:         "(id uint)",
			returnType:     "(*domain.%s, error)",
			returnTemplate: "return s.Repository.Select{{.EntityName}}(id)",
		},
		"list": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "List",
			params:         "",
			returnType:     "(*[]domain.%s, error)",
			returnTemplate: "return s.Repository.SelectAll{{.EntityName}}()",
		},
		"update": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Update",
			params:         "(id uint, newData map[string]interface{})",
			returnType:     "(*domain.%s, error)",
			returnTemplate: "return s.Repository.Update{{.EntityName}}(id, newData)",
		},
		"delete": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Delete",
			params:         "(id uint)",
			returnType:     "",
			returnTemplate: "return s.Repository.Delete{{.EntityName}}(id)",
		},
	},
	"repository": {
		"insert": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Insert",
			params:         "(%s domain.%s)",
			returnType:     "(*domain.%s, error)",
			returnTemplate: "if err := r.DB.Create(&{{.VariableName}}).Error; err != nil {\n\t\treturn nil, err\n\t}\n\tvar new{{.EntityName}} domain.{{.EntityName}}\n\tif err :=r.DB.First(&new{{.EntityName}}, {{.VariableName}}.ID).Error; err != nil {\n\t\treturn nil, err\n\t}\n\treturn &new{{.EntityName}}, nil",
		},
		"select": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Select",
			params:         "(id uint)",
			returnType:     "(*domain.%s, error)",
			returnTemplate: "var {{.VariableName}} domain.{{.EntityName}}\n\tif err := r.DB.First(&{{.VariableName}}, id).Error; err != nil {\n\t\treturn nil, err\n\t}\n\treturn &{{.VariableName}}, nil",
		},
		"selectAll": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "SelectAll",
			params:         "",
			returnType:     "(*[]domain.%s, error)",
			returnTemplate: "var {{.VariableName}} []domain.{{.EntityName}}\n\tif err := r.DB.Find(&{{.VariableName}}).Error; err != nil {\n\t\treturn nil, err\n\t}\n\treturn &{{.VariableName}}, nil",
		},
		"update": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Update",
			params:         "(id uint, newData map[string]interface{})",
			returnType:     "(*domain.%s, error)",
			returnTemplate: "var {{.VariableName}} domain.{{.EntityName}}\n\tif err := r.DB.First(&{{.VariableName}}, id).Error; err != nil {\n\t\treturn nil, err\n\t}\n\tif err := r.DB.Model(&{{.VariableName}}).Updates(newData).Error; err != nil {\n\t\treturn nil, err\n\t}\n\treturn &{{.VariableName}}, nil",
		},
		"delete": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Delete",
			params:         "(id uint)",
			returnType:     "",
			returnTemplate: "if err := r.DB.Delete(&domain.{{.EntityName}}{}, id).Error; err != nil {\n\t\treturn err\n\t}\n\treturn nil",
		},
	},
	"handler": {
		"get": {
			receiverPrefix: "func (h *Handler)",
			funcPrefix:     "Get",
			params:         "(c *fiber.Ctx)",
			returnType:     "error",
			returnTemplate: "list{{.EntityName}},err := h.Service.List{{.EntityName}}()\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{\"Message\":\"Failed to get {{.EntityName}}\", \"Error\": err.Error()})\n\t}\n\treturn c.Status(fiber.StatusOK).JSON(fiber.Map{\"Message\":\"Successfully get {{.EntityName}}\", \"Data\": list{{.EntityName}}})",
		},
		"getById": {
			receiverPrefix: "func (h *Handler)",
			funcPrefix:     "GetById",
			params:         "(c *fiber.Ctx)",
			returnType:     "error",
			returnTemplate: "id, err := c.ParamsInt(\"id\")\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"Message\": \"Invalid ID\", \"Error\": err.Error()})\n\t}\n\t{{.VariableName}}, err := h.Service.Read{{.EntityName}}(uint(id))\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{\"Message\":\"Failed to get {{.EntityName}}\", \"Error\": err.Error()})\n\t}\n\treturn c.Status(fiber.StatusOK).JSON(fiber.Map{\"Message\":\"Successfully get {{.EntityName}}\", \"Data\": {{.VariableName}}})",
		},
		"post": {
			receiverPrefix: "func (h *Handler)",
			funcPrefix:     "Post",
			params:         "(c *fiber.Ctx)",
			returnType:     "error",
			returnTemplate: "var {{.VariableName}} domain.{{.EntityName}}\n\tif err := c.BodyParser(&{{.VariableName}}); err != nil {\n\t\treturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"Message\":\"Failed to parse request body\", \"Error\": err.Error()})\n\t}\n\tnew{{.EntityName}}, err := h.Service.Create{{.EntityName}}({{.VariableName}})\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{\"Message\":\"Failed to create {{.EntityName}}\", \"Error\": err.Error()})\n\t}\n\treturn c.Status(fiber.StatusCreated).JSON(fiber.Map{\"Message\":\"Successfully created {{.EntityName}}\", \"Data\": new{{.EntityName}}})",
		},
		"put": {
			receiverPrefix: "func (h *Handler)",
			funcPrefix:     "Put",
			params:         "(c *fiber.Ctx)",
			returnType:     "error",
			returnTemplate: "id,err := c.ParamsInt(\"id\")\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"Message\": \"Invalid ID\", \"Error\": err.Error()})\n\t}\n\tvar {{.VariableName}} map[string]interface{}\n\tif err := c.BodyParser(&{{.VariableName}}); err != nil {\n\t\treturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"Message\":\"Failed to parse request body\", \"Error\": err.Error()})\n\t}\n\tupdated{{.EntityName}}, err := h.Service.Update{{.EntityName}}(uint(id),{{.VariableName}})\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{\"Message\":\"Failed to update {{.EntityName}}\", \"Error\": err.Error()})\n\t}\n\treturn c.Status(fiber.StatusOK).JSON(fiber.Map{\"Message\":\"Successfully updated {{.EntityName}}\", \"Data\": updated{{.EntityName}}})",
		},
		"delete": {
			receiverPrefix: "func (h *Handler)",
			funcPrefix:     "Delete",
			params:         "(c *fiber.Ctx)",
			returnType:     "error",
			returnTemplate: "id,err := c.ParamsInt(\"id\")\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"Message\": \"Invalid ID\", \"Error\": err.Error()})\n\t}\n\tif err := h.Service.Delete{{.EntityName}}(uint(id)); err != nil {\n\t\treturn c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{\"Message\":\"Failed to delete {{.EntityName}}\", \"Error\": err.Error()})\n\t}\n\treturn c.Status(fiber.StatusOK).JSON(fiber.Map{\"Message\":\"Successfully deleted {{.EntityName}}\"})",
		},
		"patch": {
			receiverPrefix: "func (h *Handler)",
			funcPrefix:     "Patch",
			params:         "(c *fiber.Ctx)",
			returnType:     "error",
			returnTemplate: "id,err := c.ParamsInt(\"id\")\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"Message\": \"Invalid ID\", \"Error\": err.Error()})\n\t}\n\tvar {{.VariableName}} map[string]interface{}\n\tif err := c.BodyParser(&{{.VariableName}}); err != nil {\n\t\treturn c.Status(fiber.StatusBadRequest).JSON(fiber.Map{\"Message\":\"Failed to parse request body\", \"Error\": err.Error()})\n\t}\n\tupdated{{.EntityName}}, err := h.Service.Update{{.EntityName}}(uint(id),{{.VariableName}})\n\tif err != nil {\n\t\treturn c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{\"Message\":\"Failed to update {{.EntityName}}\", \"Error\": err.Error()})\n\t}\n\treturn c.Status(fiber.StatusOK).JSON(fiber.Map{\"Message\":\"Successfully updated {{.EntityName}}\", \"Data\": updated{{.EntityName}}})",
		},
	},
}

func GenerateSignature(layer, funcNameLayer, entityName string) string {
	config, ok := templates[layer][funcNameLayer]
	if !ok {
		return "not found"
	}

	lowerEntity := strings.ToLower(entityName)

	// Formatear parámetros
	formattedParams := fmt.Sprintf(config.params, lowerEntity, entityName)
	if strings.Contains(formattedParams, "%!") {
		formattedParams = config.params
	}
	if config.params == "" {
		formattedParams = "()"
	}
	//formatear el string de retorno
	stringReturn := fmt.Sprintf(config.returnType, entityName)
	if config.returnType == "" {
		stringReturn = "error"
	}

	stringReturnNew, err := GenerateReturnParams(config.returnTemplate, map[string]interface{}{
		"EntityName":   entityName,
		"VariableName": lowerEntity,
	})
	if err != nil {
		return err.Error()
	}

	// Construir firma de la función
	return fmt.Sprintf(
		"\n%s %s%s%s %s {\t\n\t%s\n}\n",
		config.receiverPrefix,
		config.funcPrefix,
		entityName,
		formattedParams,
		stringReturn,
		stringReturnNew,
	)
}

func GenerateReturnParams(templateString string, data map[string]interface{}) (string, error) {
	var renderedTemplate bytes.Buffer
	tmpl, err := template.New("returnParams").Parse(templateString)
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(&renderedTemplate, data); err != nil {
		return "", err
	}
	return renderedTemplate.String(), nil
}

func GenerateHandler(layer, funcNameLayer, entityName string) string {
	config, ok := templates[layer][funcNameLayer]
	if !ok {
		return "not found"
	}
	returnParams, err := GenerateReturnParams(config.returnTemplate, map[string]interface{}{
		"EntityName":   entityName,
		"VariableName": strings.ToLower(entityName),
	})
	if err != nil {
		return err.Error()
	}

	res := fmt.Sprintf("%s %s%s%s %s {\t\n\t%s\n}\n",
		config.receiverPrefix,
		config.funcPrefix,
		entityName,
		config.params,
		config.returnType,
		returnParams,
	)
	return res
}
