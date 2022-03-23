package main

import (
	"fmt"
	"strconv"
)

func Pourcentage(sbody string) float64 {
	var pourcentage float64
	diviseur := ""
	dividende := ""
	condition := false
	for i := 0; i <= len(sbody)-1; i++ {
		if rune(sbody[i]) == 32 {
			condition = true
		} else if condition == true {
			dividende = dividende + string(sbody[i])
		} else if condition == false {
			diviseur = diviseur + string(sbody[i])
		}
	}
	int_diviseur, err := strconv.ParseFloat(diviseur, 64)
	int_dividende, err := strconv.ParseFloat(dividende, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	pourcentage = ((1 - int_dividende/int_diviseur) * 100) * -1
	return pourcentage
}
