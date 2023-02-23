package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func reciptHandler(c *fiber.Ctx) error {

	recipt, err := getRecipt()
	if err != nil {
		log.Println("error in getRecipt", err)
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(fiber.Map{"success": true, "data": recipt})
}

func singleRecpitHandler(c *fiber.Ctx) error {

	return c.Status(500).JSON(fiber.Map{"success": true, "data": nil})
}

func createServer() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/persons", reciptHandler)
	app.Get("/persons/:id", singleRecpitHandler)

	log.Fatal(app.Listen(":3000"))
}
