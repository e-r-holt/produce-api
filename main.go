package main

import (
	"fmt"
	"log"

	"github.com/e-r-holt/produce-api/db"
	store "github.com/e-r-holt/produce-api/db" //"database" operations

	"github.com/gofiber/fiber/v2" //API framework
)

func main() {
	db := db.Database()
	app := fiber.New()

	// GET w/ optional parameter
	app.Get("/:produce_code?", func(c *fiber.Ctx) error {
		//if param given
		code := c.Params("produce_code")
		if code != "" {
			data, err := db.ReadOne(code)
			if err != nil {
				fmt.Println(err)
				return c.SendString("Couldn't find that one")
			} else {
				return c.JSON(data)
			}
		} else { //if no param
			return c.JSON(db)
		}
	})

	app.Post("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		// Get raw body from POST request:
		new := new(store.Produce)
		if err := c.BodyParser(new); err != nil {
			return err
		}

		// if db.IsDuplicate(new.Code) != true {

		// }
		return c.JSON(new) // []byte("user=john")
	})
	log.Fatal(app.Listen(":3000"))
}
