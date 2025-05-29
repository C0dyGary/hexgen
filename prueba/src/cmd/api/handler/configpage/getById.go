package configpage

import (
	"github.com/C0dyGary/gift/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
)
func (h *Handler) GetByIdConfigPage(c *fiber.Ctx) error {	
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Invalid ID", "Error": err.Error()})
	}
	configpage, err := h.Service.ReadConfigPage(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Failed to get ConfigPage", "Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message":"Successfully get ConfigPage", "Data": configpage})
}
