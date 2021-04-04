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

	//expect list of json objects to add to DB
	app.Post("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		// Get raw body from POST request:
		new := new(store.ProduceSlice)
		if err := c.BodyParser(new); err != nil {
			return err
		} else {
			//dupe checking
			dupe := false
			for _, v := range *new {
				if dupe = db.IsDuplicate(v.Code); dupe {
					c.SendString("No duplicates allowed: " + v.Code)
					return c.SendStatus(409)
				}
			}
			if !dupe { //if no dupes, go create
				c.SendString("Start creating")
				return c.SendStatus(201)
			} else {
				c.SendString("We shouldn't have gotten here!")
				return c.SendStatus(500)
			}
		}

		// if !db.IsDuplicate(new.Code) {
		// 	return c.JSON(db.CreateOne(*new))
		// } else {
		// 	return c.SendString("Can't create duplicates!")
		// }
	})
	log.Fatal(app.Listen(":3000"))
}
