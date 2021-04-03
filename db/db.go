//Produce database with CR-D operations
package db

import "errors"

//Structure of db records
type Produce struct {
	Code  string
	Name  string
	Price float64
}
type ProduceSlice []Produce

//return one record
//arg == Produce Code
func (ps ProduceSlice) ReadOne(code string) (*Produce, error) {
	for _, v := range ps {
		// fmt.Printf("Index = %d, Value = %s", i, v.Name)
		if v.Code == code {
			// fmt.Println()
			// fmt.Println("Found!")
			return &v, nil
		}
	}
	return nil, errors.New("can't find the produce code")
}

//return all records
// func (ps *ProduceSlice) ReadAll() {
// }

//add a new record
func (ps ProduceSlice) CreateOne(new Produce) ProduceSlice {
	ps = append(ps, new)
	return ps
}

//create many records
func (ps ProduceSlice) CreateMany(new []Produce) ProduceSlice {
	for _, v := range new {
		ps = append(ps, v)
	}
	return ps
}

//delete one record, returns modified slice
//if no change, returns given slice
func (ps ProduceSlice) DeleteOne(code string) ProduceSlice {

	for i, v := range ps {
		if v.Code == code {
			ps[i] = ps[len(ps)-1]
			//no need to put ps[i] at the end, since it will be discarded
			return ps[:len(ps)-1]
		}
	}
	return ps
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
