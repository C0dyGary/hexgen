package handlers

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func GenMkdir(c *cli.Context) error {
	directories := []string{"src/pkg/domain", "src/pkg/port", "src/pkg/repository", "src/pkg/service", "src/cmd/api/handler", "src/cmd/api/routes"}
	for _, dir := range directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	fmt.Println("Estructura creada con éxito.")
	return nil
}
