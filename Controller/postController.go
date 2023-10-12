package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/habibbushira/goblog/database"
	"github.com/habibbushira/goblog/models"
	"github.com/habibbushira/goblog/util"
	"gorm.io/gorm"
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
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64

	var blogs []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&blogs)
	database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": blogs,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func Post(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blog models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blog)

	if blog.Id == 0 {
		c.Status(fiber.StatusNoContent)
		return nil
	}

	return c.JSON(fiber.Map{
		"data": blog,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := &models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)

	return c.JSON(fiber.Map{
		"message": "Blog updated successfully",
	})
}

func MyPosts(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var blogs []models.Blog
	database.DB.Model(&blogs).Where("user_id=?", id).Preload("User").Find(&blogs)

	return c.JSON(blogs)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blog = models.Blog{
		Id: uint(id),
	}

	deleteBlog := database.DB.Delete(&blog)
	if errors.Is(deleteBlog.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps!, record not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Blog removed successfully",
	})
}
