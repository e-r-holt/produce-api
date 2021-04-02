package db

type Produce struct {
	Name  string
	Code string
	Price float64
}

func Database() (db map[string]Produce) {
	db = make(map[string]Produce)
	db["A12T-4GH7-QPL9-3N4M"] = Produce{"Lettuce", 3.46}
	db["E5T6-9UI3-TH15-QR88"] = Produce{"Peach", 2.99}
	db["YRT6-72AS-K736-L4AR"] = Produce{"Green Pepper", 0.79}
	db["TQ4C-VV6T-75ZX-1RMR"] = Produce{"Gala Apple", 3.59}
	return
}
