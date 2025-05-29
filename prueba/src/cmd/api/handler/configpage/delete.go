package configpage

import (
	"github.com/C0dyGary/gift/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
)
func (h *Handler) DeleteConfigPage(c *fiber.Ctx) error {	
	id,err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Invalid ID", "Error": err.Error()})
	}
	if err := h.Service.DeleteConfigPage(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Failed to delete ConfigPage", "Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message":"Successfully deleted ConfigPage"})
}
