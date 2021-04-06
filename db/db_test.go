package db

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	data := Database()
	res := make(chan ProduceSlice)
	err := make(chan string)

	//good case
	go data.ReadOne("A12T-4GH7-QPL9-3N4M", res, err)
	select {
	case some := <-res:
		t.Log(some)
	case thing := <-err:
		t.Error(thing)
	}

	//bad case
	go data.ReadOne("bad", res, err)
	select {
	case some := <-res:
		t.Error(some)
	case thing := <-err:
		t.Log(thing)
	}
}

// func TestCreate(t *testing.T) {
// 	data := Database()

// 	new := Produce{"asdf", "asdf", 3.14}
// 	beforeLen := len(data)
// 	res := make(chan ProduceSlice)
// 	err := make(chan string)
// 	go data.CreateOne(new, res, err)

// 	go data.ReadOne("bad", res, err)
// 	select {
// 	case some := <-res:
// 		t.Error(some)
// 	case thing := <-err:
// 		t.Log(thing)
// 	}
// 	afterLen := len(data)
// 	if beforeLen >= afterLen {
// 		t.Error("did not properly append")
// 	}
// }

func TestCreateMany(t *testing.T) {

	data := Database()
	beforeLen := len(data)
	new := []Produce{
		{Code: "A12T-4GH7-QPL9-3N4M", Name: "Lettuce", Price: 3.46},
		{Code: "E5T6-9UI3-TH15-QR88", Name: "Peach", Price: 2.99},
		{Code: "YRT6-72AS-K736-L4AR", Name: "Green Pepper", Price: 0.79},
		{Code: "TQ4C-VV6T-75ZX-1RMR", Name: "Gala Apple", Price: 3.59},
	}
	appendLen := len(new)
	res := make(chan ProduceSlice)
	go data.CreateMany(new, res)

	select {
	case data := <-res:
		afterLen := len(data)
		if afterLen != (beforeLen + appendLen) {
			t.Error("did not append all records")
		}
	}
}

func TestDeleteOne(t *testing.T) {

	data := Database()
	beforeLen := len(data)
	deleteme := "A12T-4GH7-QPL9-3N4M"

	data, err := data.DeleteOne(deleteme)
	afterLen := len(data)
	//if no err, delete must have happened
	if err != nil {
		fmt.Println(err)
		afterLen := len(data)
		if afterLen <= beforeLen {
			t.Error("did not delete all records")
		}
	} else if afterLen < beforeLen {
		t.Log("Successfully deleted one")
	} else {
		t.Error("Failed in deletion")
	}

}

func TestBadDelete(t *testing.T) {
	data := Database()
	beforeLen := len(data)
	deleteme := "asdf"

	data, err := data.DeleteOne(deleteme)
	afterLen := len(data)
	//if no err, delete must have happened
	if err != nil {
		afterLen := len(data)
		if afterLen == beforeLen {
			t.Log("PASS: code does not exist to delete")
		}
	} else if afterLen < beforeLen {
		t.Log("")
	}

}

func TestIsDuplicate(t *testing.T) {
	data := Database()
	codes := [2]string{"A12T-4GH7-QPL9-3N4M", "bad"}

	//should return true
	isDup := data.IsDuplicate(codes[0])
	if isDup == true {
		t.Log("PASS: found duplicate")
	} else {
		t.Error("failed to identify duplicate")
	}

	//should return false
	isDup = data.IsDuplicate(codes[1])
	if isDup == false {
		t.Log("PASS: not a dupe")
	} else {
		t.Error("returned false for code that dne")
	}
}
