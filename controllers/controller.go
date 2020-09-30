package controllers

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

// Ping handler
func Ping(c *fiber.Ctx) error  {
	c.Status(http.StatusOK)
	c.SendString("PONG")

	return nil
}

// ChainHandler is used to move forward the request
func ChainHandler(c *fiber.Ctx) error  {
	return c.Next()
}

// UploadFile handler 
func UploadFile(c *fiber.Ctx) error  {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatalln("Error occured while file upload")
		c.SendStatus(http.StatusBadRequest)
	}

	log.Println(file)
	return c.SendString("Upload file is not yet implemented")
}