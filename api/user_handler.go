package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redmicelles/Hotel-Reservation/types"
)

func HandleListUsers(c *fiber.Ctx) error {
	u := types.User{
		ID:        "",
		FirstName: "Nathan",
		LastName:  "Daniel",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Sola"})
}
