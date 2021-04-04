package main

import (
	"log"

	"github.com/e-r-holt/produce-api/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := db.Database()
	app := fiber.New()

	app.Get("/:produce_code", func(c *fiber.Ctx) error {
		data, err := db.ReadOne("A12T-4GH7-QPL9-3N4M")
		if err != nil {
			return c.JSON(data)
		} else {
			return c.SendString("Error")
		}
	})
	log.Fatal(app.Listen(":3000"))
}
