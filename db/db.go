package db

type Produce struct {
	Code  string
	Name  string
	Price float64
}

func Database() (db []Produce) {
	db = []Produce{
		{"A12T-4GH7-QPL9-3N4M", "Lettuce", 3.46},
		{"E5T6-9UI3-TH15-QR88", "Peach", 2.99},
		{"YRT6-72AS-K736-L4AR", "Green Pepper", 0.79},
		{"TQ4C-VV6T-75ZX-1RMR", "Gala Apple", 3.59},
	}
	return
}
