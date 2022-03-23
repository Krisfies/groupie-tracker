package main

func Display(body string) string {
	sbody := ""
	virgule := 0
	accolade := 0
	correct := false
	affibis := true
	for i := 0; i <= len(string(body))/2; i++ {
		if rune(body[i]) == 44 {
			correct = true
			virgule = virgule + 1
		} else if rune(body[i]) == 93 && accolade < 2 {
			if affibis == true {
				sbody = sbody + " "
				// Prix actuel
				affibis = false
				correct = false
				accolade = accolade + 1
			} else {
				// Prix 24h
				correct = false
				accolade = accolade + 1
			}
		} else if correct == true && virgule != 2 {
			sbody = sbody + string(body[i])
		} else if accolade == 2 {
			return sbody
		}
	}
	return ""
}
