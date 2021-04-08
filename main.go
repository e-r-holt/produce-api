package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2" //API framework
)

func appSetup() *fiber.App {
	db := Database()
	app := fiber.New()

	// GET w/ optional parameter
	app.Get("/:produce_code?", func(c *fiber.Ctx) error {
		res := make(chan ProduceSlice)
		err := make(chan string)
		//if param given
		code := c.Params("produce_code")
		if code != "" {
			go db.ReadOne(code, res, err)
			select {
			case record := <-res:
				return c.JSON(record)
			case error := <-err:
				return c.SendString(error)
			}
		} else { //if no param
			return c.JSON(db)
		}
	})

	//expect list of json objects to add to DB
	app.Post("/", func(c *fiber.Ctx) error {
		res := make(chan ProduceSlice)
		err := make(chan string)
		c.Accepts("application/json")

		// Get raw body from POST request:
		new := new(ProduceSlice)
		if errorStr := c.BodyParser(new); errorStr != nil {
			fmt.Println(errorStr)
			return errorStr
		} else {
			// dupe checking
			var dupes ProduceSlice
			for _, v := range *new {
				go db.ReadOne(v.Code, res, err)
			}
			for range *new {
				select {
				case dupRec := <-res:
					dupes = append(dupes, dupRec...)
					// fmt.Println("found a dupe")
				case <-err: //err means not dupe
					// fmt.Println("Not a dupe")
				}
			}

			if len(dupes) == 0 { //if no dupes, go create
				for _, v := range *new {
					go func(v Produce) {
						db = append(db, v)
					}(v)
				}

				c.SendString("Added records to the db")
				return c.SendStatus(201)
			} else {
				c.JSON(dupes)
				return c.SendStatus(409)
			}
		}
	})

	return app

}
func main() {
	app := appSetup()
	log.Fatal(app.Listen(":3000"))
}
