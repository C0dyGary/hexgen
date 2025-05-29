package configpage

import (
	"github.com/C0dyGary/gift/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
)
func (h *Handler) GetConfigPage(c *fiber.Ctx) error {	
	listConfigPage,err := h.Service.ListConfigPage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Failed to get ConfigPage", "Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message":"Successfully get ConfigPage", "Data": listConfigPage})
}
