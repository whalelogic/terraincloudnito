package handlers

import (
	"github.com/gofiber/fiber/v2"
)


func DownloadResume(c *fiber.Ctx) error {
	return c.Download("./static/files/resume.pdf", "Keith_Thomson_Resume.pdf")
}
