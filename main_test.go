// myfile_test.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	//API framework
	utils "github.com/gofiber/utils"
)

// go test -run -v Test_Handler
func TestGetAll(t *testing.T) {
	app := appSetup()

	// test get all
	url := "/"
	resp, err := app.Test(httptest.NewRequest("GET", url, nil))
	utils.AssertEqual(t, nil, err, "app.Test")

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		t.Error(readErr)
	}

	expect := []Produce{
		{"A12T-4GH7-QPL9-3N4M", "Lettuce", 3.46},
		{"E5T6-9UI3-TH15-QR88", "Peach", 2.99},
		{"YRT6-72AS-K736-L4AR", "Green Pepper", 0.79},
		{"TQ4C-VV6T-75ZX-1RMR", "Gala Apple", 3.59},
	}
	var respProduce []Produce
	jsonErr := json.Unmarshal(body, &respProduce)
	if jsonErr != nil {
		log.Printf("error decoding response: %v", jsonErr)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", body)
	}
	utils.AssertEqual(t, expect, respProduce, "Payload v Response")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	// ################################################################
}

func TestGetValid(t *testing.T) {
	app := appSetup()

	//get one good record
	var expect = []Produce{
		{
			Code:  "E5T6-9UI3-TH15-QR88",
			Name:  "Peach",
			Price: 2.99,
		},
	}
	url := "/" + expect[0].Code
	resp, err := app.Test(httptest.NewRequest("GET", url, nil))
	utils.AssertEqual(t, nil, err, "app.Test")

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		t.Error(readErr)
	}

	respProduce := []Produce{}
	jsonErr := json.Unmarshal(body, &respProduce)
	if jsonErr != nil {
		log.Printf("error decoding response: %v", jsonErr)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", body)
	}
	utils.AssertEqual(t, expect, respProduce, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

func TestGetInalid(t *testing.T) {
	app := appSetup()

	url := "/asdf"
	resp, err := app.Test(httptest.NewRequest("GET", url, nil))
	utils.AssertEqual(t, 404, resp.StatusCode, "Status code")
	utils.AssertEqual(t, nil, err, "app.Test")
}

func TestCreateOne(t *testing.T) {
	app := appSetup()

	url := "/"
	reqPayload, err := json.Marshal([]Produce{
		{
			Code:  "some-ting-real-cool",
			Name:  "a",
			Price: 5.46,
		},
	})
	// fmt.Println(payload)
	if err != nil {
		t.Error(err)
	}
	body := string(reqPayload)
	// fmt.Println(body)
	req := httptest.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	utils.AssertEqual(t, 201, resp.StatusCode, "Status code")
	utils.AssertEqual(t, nil, err, "app.Test")

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	returned, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		t.Error(readErr)
	}

	respProduce := []Produce{}
	// val, err := strconv.Unquote(string(returned))
	// if err != nil {
	// 	t.Error(err)
	// }
	jsonErr := json.Unmarshal(returned, &respProduce)
	if jsonErr != nil {
		log.Printf("error decoding response: %v", jsonErr)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", returned)
	}
	utils.AssertEqual(t, reqPayload, returned, "Payload v Response")
}

func TestCreateMany(t *testing.T) {
	app := appSetup()

	url := "/"
	reqPayload, err := json.Marshal([]Produce{
		{
			Code:  "some-ting-real-cool",
			Name:  "a",
			Price: 5.46,
		},
		{
			Code:  "some-tiag-real-cool",
			Name:  "b",
			Price: 5.46,
		},
		{
			Code:  "aome-ting-real-cool",
			Name:  "c",
			Price: 5.46,
		},
		{
			Code:  "some-ting-real-aool",
			Name:  "d",
			Price: 5.46,
		},
		{
			Code:  "some-ting-real-coo",
			Name:  "e",
			Price: 5.46,
		},
		{
			Code:  "some-tingreal-cool",
			Name:  "f",
			Price: 5.46,
		},
		{
			Code:  "some-ting-realcool",
			Name:  "g",
			Price: 5.46,
		},
		{
			Code:  "some-ting-real-cool",
			Name:  "h",
			Price: 5.46,
		},
	})
	// fmt.Println(payload)
	if err != nil {
		t.Error(err)
	}
	body := string(reqPayload)
	// fmt.Println(body)
	req := httptest.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	utils.AssertEqual(t, 201, resp.StatusCode, "Status code")
	utils.AssertEqual(t, nil, err, "app.Test")

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	returned, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		t.Error(readErr)
	}

	respProduce := []Produce{}
	// val, err := strconv.Unquote(string(returned))
	// if err != nil {
	// 	t.Error(err)
	// }
	jsonErr := json.Unmarshal(returned, &respProduce)
	if jsonErr != nil {
		log.Printf("error decoding response: %v", jsonErr)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", returned)
	}
	utils.AssertEqual(t, reqPayload, returned, "Payload v Response")
}

func TestDelValid(t *testing.T) {
	app := appSetup()

	delMe := []Produce{{"A12T-4GH7-QPL9-3N4M", "Lettuce", 3.46}}
	//del one good record

	url := "/" + delMe[0].Code
	resp, err := app.Test(httptest.NewRequest("DELETE", url, nil))
	utils.AssertEqual(t, nil, err, "app.Test")

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		t.Error(readErr)
	}

	respProduce := []Produce{}
	jsonErr := json.Unmarshal(body, &respProduce)
	if jsonErr != nil {
		log.Printf("error decoding response: %v", jsonErr)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", body)
	}
	expect := []Produce{
		{"TQ4C-VV6T-75ZX-1RMR", "Gala Apple", 3.59},
		{"E5T6-9UI3-TH15-QR88", "Peach", 2.99},
		{"YRT6-72AS-K736-L4AR", "Green Pepper", 0.79},
	}

	for _, v := range expect {
		if v.Code == delMe[0].Code {
			t.Error("Did not delete")
		}
	}
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

func TestDelInalid(t *testing.T) {
	app := appSetup()

	//del one good record

	url := "/asdf"
	resp, err := app.Test(httptest.NewRequest("DELETE", url, nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 404, resp.StatusCode, "Status code")
}
