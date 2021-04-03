package db

import "errors"

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
func (ps *ProduceSlice) ReadAll() {

}

//create one record
func (ps *ProduceSlice) CreateOne() {

}

//create many records
func (ps *ProduceSlice) CreateMany() {

}

//delete one record
func (ps *ProduceSlice) DeleteOne() {

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
