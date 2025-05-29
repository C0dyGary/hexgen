package configpage

import (
	"github.com/C0dyGary/gift/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
)
func (h *Handler) PatchConfigPage(c *fiber.Ctx) error {	
	id,err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Invalid ID", "Error": err.Error()})
	}
	var configpage map[string]interface{}
	if err := c.BodyParser(&configpage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message":"Failed to parse request body", "Error": err.Error()})
	}
	updatedConfigPage, err := h.Service.UpdateConfigPage(uint(id),configpage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Failed to update ConfigPage", "Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message":"Successfully updated ConfigPage", "Data": updatedConfigPage})
}
