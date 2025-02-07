package template

import (
	"fmt"
	"strings"
)

type Config struct {
	receiverPrefix string
	funcPrefix     string
	params         string
	returnType     string
}

var templates = map[string]map[string]Config{
	"service": {
		"create": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Create",
			params:         "(%s domain.%s)",
			returnType:     "(*domain.%s, error)",
		},
		"read": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Read",
			params:         "(id uint)",
			returnType:     "(*domain.%s, error)",
		},
		"list": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "List",
			params:         "",
			returnType:     "(*[]domain.%s, error)",
		},
		"update": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Update",
			params:         "(id uint, newData map[string]interface{})",
			returnType:     "(*domain.%s, error)",
		},
		"delete": {
			receiverPrefix: "func (s *Service)",
			funcPrefix:     "Delete",
			params:         "(id uint)",
			returnType:     "",
		},
	},
	"repository": {
		"insert": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Insert",
			params:         "(%s domain.%s)",
			returnType:     "(*domain.%s, error)",
		},
		"select": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Select",
			params:         "(id uint)",
			returnType:     "(*domain.%s, error)",
		},
		"selectAll": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "SelectAll",
			params:         "",
			returnType:     "(*[]domain.%s, error)",
		},
		"update": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Update",
			params:         "(id uint, newData map[string]interface{})",
			returnType:     "(*domain.%s, error)",
		},
		"delete": {
			receiverPrefix: "func (r *Repository)",
			funcPrefix:     "Delete",
			params:         "(id uint)",
			returnType:     "",
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

	// Construir firma de la función
	return fmt.Sprintf(
		"\n%s %s%s%s %s {\t\n\treturn nil, nil // TODO: Implementar la lógica\n}\n",
		config.receiverPrefix,
		config.funcPrefix,
		entityName,
		formattedParams,
		stringReturn,
	)
}
