package cmd

import (
	"fmt"
	"github.com/C0dyGary/hexgen/src/cmd/handlers"
	"github.com/urfave/cli/v2"
	"os"
)

func StartCli() {

	app := &cli.App{
		Name:  "hexgen",
		Usage: "Generates a hexagonal architecture structure en Go",
		Commands: []*cli.Command{
			{
				Name:   "crt",
				Usage:  "Crea la estructura base de un proyecto en Go",
				Action: handlers.GenMkdir,
			},
			{
				Name:  "str",
				Usage: "Genera una entidad",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true, Usage: "Nombre de la estructura"},
					&cli.StringSliceFlag{Name: "fields", Aliases: []string{"f"}, Usage: "Lista de campos de la estructura formato: Name:type"},
				},
				Action: handlers.GenEntity,
			},
			{
				Name:  "gen",
				Usage: "Genera el puerto, service, repository de la entidad ya creada",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true, Usage: "Nombre de la estructura en formato PascalCase Ej:'User','Role'"},
					&cli.StringFlag{Name: "project", Aliases: []string{"p"}, Required: true, Usage: "Nombre del proyecto"}},
				Action: handlers.GenTemplate,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
