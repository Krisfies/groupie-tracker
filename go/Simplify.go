package main

import "strconv"

func Simplify(nbr float64) float64 {
	snbr := strconv.FormatFloat(nbr, 'f', 1, 64)
	new_nbr, err := strconv.ParseFloat(snbr, 64)
	if err != nil {
		return 0.0
	}
	return new_nbr
}
