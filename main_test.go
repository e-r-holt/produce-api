// myfile_test.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	fiber "github.com/gofiber/fiber/v2" //API framework
	utils "github.com/gofiber/utils"
)

// go test -run -v Test_Handler
func Test_Get(t *testing.T) {
	app := fiber.New()

	//good record
	var expect = []Produce{
		{
			Code:  "E5T6-9UI3-TH15-QR88",
			Name:  "Peach",
			Price: 2.99,
		},
	}
	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	utils.AssertEqual(t, nil, err, "app.Test")

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		t.Error(readErr)
	}

	respProduce := Produce{}
	// jsonErr := json.Unmarshal(body, &respProduce)
	// if jsonErr != nil {
	// 	t.Error(jsonErr)
	// }
	jsonErr := json.Unmarshal(body, respProduce)
	if jsonErr != nil {
		log.Printf("error decoding response: %v", jsonErr)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", body)
	}
	// utils.AssertEqual(t, expect, respProduce, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
