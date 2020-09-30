package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/undriku/go-fiber-upload/controllers"
)

func init() {
	log.Println("All the initialization should be done here...")
}

func main() {
	server := fiber.New()
	server.Use(logger.New())

	server.Get("/ping", controllers.Ping)

	api := server.Group("/api",controllers.ChainHandler)

	v1 := api.Group("/v1", controllers.ChainHandler)
	v1.Post("/upload", controllers.UploadFile)

	// Start the server
	server.Listen(":50000")
}
