package db

import (
	"testing"
)

// func TestDatabase(t *testing.T) {
// 	data := Database()

// 	for i := 0; i < len(data); i++ {
// 		fmt.Println("Code", data[i].Code)
// 		fmt.Println("Name", data[i].Name)
// 		fmt.Println("Price $", data[i].Price)
// 		fmt.Println()
// 	}

// }

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
