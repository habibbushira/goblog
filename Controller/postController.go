package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/habibbushira/goblog/database"
	"github.com/habibbushira/goblog/models"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPost models.Blog
	if err := c.BodyParser(&blogPost); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Blog created successfully",
	})
}

func Posts(c *fiber.Ctx) error {
	return nil
}
