package db

import (
	"testing"
	"fmt"
)


func TestDatabase(t *testing.T) {
	data := Database()

	for i:=0; i<len(data); i++ {
		fmt.Println("Code", data[i].Code)
		fmt.Println("Name", data[i].Name)
		fmt.Println("Price $", data[i].Price)
		fmt.Println()
	}
	
}