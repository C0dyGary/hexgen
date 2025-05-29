package configpage

import (
	"github.com/C0dyGary/gift/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
)
func (h *Handler) PostConfigPage(c *fiber.Ctx) error {	
	var configpage domain.ConfigPage
	if err := c.BodyParser(&configpage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message":"Failed to parse request body", "Error": err.Error()})
	}
	newConfigPage, err := h.Service.CreateConfigPage(configpage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Failed to create ConfigPage", "Error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Message":"Successfully created ConfigPage", "Data": newConfigPage})
}
