package db

import (
	"testing"
)

func TestGoodRead(t *testing.T) {
	data := Database()

	pro, err := data.ReadOne("A12T-4GH7-QPL9-3N4M")
	if err != nil {
		t.Error("raised error for good read")
	}
	if pro.Name != "Lettuce" {
		t.Error("pulled the wrong produce")
	}

}

func TestBadRead(t *testing.T) {
	data := Database()

	_, err := data.ReadOne("foo")
	if err == nil {
		t.Error("did not error on bad code")
	}

}

func TestCreate(t *testing.T) {
	data := Database()

	new := Produce{"asdf", "asdf", 3.14}
	beforeLen := len(data)
	data = data.CreateOne(new)

	afterLen := len(data)
	if beforeLen >= afterLen {
		t.Error("did not properly append")
	}
}

func TestCreateMany(t *testing.T) {

	data := Database()
	beforeLen := len(data)
	new := []Produce{
		Produce{Code: "A12T-4GH7-QPL9-3N4M", Name: "Lettuce", Price: 3.46},
		Produce{Code: "E5T6-9UI3-TH15-QR88", Name: "Peach", Price: 2.99},
		Produce{Code: "YRT6-72AS-K736-L4AR", Name: "Green Pepper", Price: 0.79},
		Produce{Code: "TQ4C-VV6T-75ZX-1RMR", Name: "Gala Apple", Price: 3.59},
	}
	appendLen := len(new)
	data = data.CreateMany(new)
	afterLen := len(data)

	if afterLen != (beforeLen + appendLen) {
		t.Error("did not append all records")
	}
}
