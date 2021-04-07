//Produce database with CR-D operations
package main

import (
	"errors"
)

//Structure of db records
type Produce struct {
	Code  string  `json:"Produce Code" validate: "required,alphanum"`
	Name  string  `json:"Name" validate: "required,alphanum"`
	Price float64 `json:"Unit Price" validate: "required`
}

type ProduceSlice []Produce

//return one record
//arg == Produce Code
func (ps ProduceSlice) ReadOne(code string, res chan ProduceSlice, err chan string) {
	for _, v := range ps {
		// fmt.Printf("Index = %d, Value = %s", i, v.Name)
		if v.Code == code {
			get := []Produce{v}
			res <- get
			return
		}
	}
	err <- code
}

//return all records
// func (ps *ProduceSlice) ReadAll() {
// }

//add a new record
func (ps ProduceSlice) CreateOne(data *ProduceSlice, new Produce, res chan ProduceSlice) {
	ps = append(*data, new)
}

//create many records
func (ps ProduceSlice) CreateMany(new []Produce, res chan ProduceSlice) {
	ps = append(ps, new...)
	res <- ps
}

//delete one record, returns modified slice
//if no change, returns given slice
func (ps ProduceSlice) DeleteOne(code string) (ProduceSlice, error) {

	for i, v := range ps {
		if v.Code == code {
			ps[i] = ps[len(ps)-1]
			//no need to put ps[i] at the end, since it will be discarded
			return ps[:len(ps)-1], nil
		}
	}
	return ps, errors.New("given produce code not in db")
}

//isDuplicate: return bool
// true == duplicate produce code
// false == produce code is new
func (ps ProduceSlice) IsDuplicate(code string) bool {
	isDup := false
	for _, v := range ps {
		if v.Code == code {
			isDup = true
		}
	}
	return isDup
}

//initialize database for produce API
func Database() (db ProduceSlice) {
	db = ProduceSlice{
		{"A12T-4GH7-QPL9-3N4M", "Lettuce", 3.46},
		{"E5T6-9UI3-TH15-QR88", "Peach", 2.99},
		{"YRT6-72AS-K736-L4AR", "Green Pepper", 0.79},
		{"TQ4C-VV6T-75ZX-1RMR", "Gala Apple", 3.59},
	}
	return
}
