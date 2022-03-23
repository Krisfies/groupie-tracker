package main

func Price_Convert(bitcoin float64, eur bool, price float64) float64 {
	if eur == true {
		price = price * 0.91
	} else {
		price = price * bitcoin
	}
	return price
}
