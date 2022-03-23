package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type crypto struct {
	Price          string
	Interval       float64
	Image          string
	Tendance       string
	Image2         string
	Price_btc      float64
	Price_tendance float64
	Name           string
	User           string
	Votes_P        string
	Votes_N        string
}

func Pair(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false

	}
}

func Simplify(nbr float64) float64 {
	snbr := strconv.FormatFloat(nbr, 'f', 1, 64)
	new_nbr, err := strconv.ParseFloat(snbr, 64)
	if err != nil {
		return 0.0
	}
	return new_nbr
}

func Price_Convert(bitcoin float64, eur bool, price float64) float64 {
	if eur == true {
		price = price * 0.91
	} else {
		price = price * bitcoin
	}
	return price
}

// Le principal outil que nous avons crées afin de trier les structures de l'api
func tracker(arg string, body string, Is_image bool, Is_number bool) string {
	// On a besoin comme argument obligatoirement de arg (le nom de la structure que l'on veux) ainsi que de body,
	// l'ensemble des structures de l'api  transformé sous forme de string. Le reste est optionnel sous pour les
	// nombres et les images.
	var affichage string                // C'est le string du rendu final
	var affichage_image string          // C'est le string du rendu final si le rendu est une image
	var affichage_value string          // C'est le string du rendu final si le rendu est un nombre
	break_ := 0                         // L'index pour stopper le programme
	pof := 1                            // La condition "pile ou face" pour avoir un index pair ou impair
	verif := ""                         // Le string de qui prend la valeur des noms de structures et vérifie que c'est le bon
	for i := 1; i <= len(body)-1; i++ { // Boucle allant de 1 a la fin de la structure
		if rune(body[i]) == '"' { // C'est la condition qui permet de commencer à verifier le nom
			pof = pof + 1 // Ajoute 1 à pof
		} else if Pair(pof) == true { // Condition si pof est pair
			verif = verif + string(body[i]) // Ecris le nom des structs dans verif
		} else if verif == arg { // C'est la condition si le nom de la struct rst le bon nom
			for y := i - 1; y <= len(body)-1; y++ { // Boucle allant du point oû l'on s'est stopper à la fin
				if rune(body[y]) == '"' { // C'est la condition qui detecte les guillemets
					break_ = break_ + 1 // Ajoute 1 au break
				} else if rune(body[y]) == '}' || rune(body[y]) == ',' { // Condition qui arrete la recherche
					for i := 1; i <= len(affichage)-1; i++ { // Boucle allant de 1 à la fin du mot
						affichage_value = affichage_value + string(affichage[i]) // Ecris seulement les chiffres
					}
					return affichage_value // Renvoie le resultat (nombre)
				} else if break_ == 3 { // Condition qui arrete aussi la recherche
					if Is_image == true { // Condition si on veux une image
						for i := 1; i <= len(affichage)-1; i++ { // Boucle allant de 1 à la fin du mot
							affichage_image = affichage_image + string(affichage[i]) // Recupere l'adresse de l'image
						}
						return affichage_image // Renvoie le resultat (image)
					} else {
						return affichage // Renvoie le resultat
					}
				} else {
					affichage = affichage + string(body[y]) // Ecris le resultat final si on ne stoppe pas la recherche
				}
			}
		} else if Pair(pof) == false { // Si pof est impair
			verif = "" // On  réinitialise le string verif
		}
	}
	return "" // Le return par défaut
}

func affichage(body string) string {
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

func traitement(url string) string {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Cookie", "__cf_bm=qfyI8Hsb8DV1Rrw16Ty.iVFlQ7JrbQqUii65A0_VxEY-1647267811-0-AbAU3Pao1XQoNIodtQn0iGOIgOKpPyWOTUQJagpQ8CABKjIqKiZXeMl4bJryS/TWZrDuZmNi5ayzefcpMhQeHjI=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

func main() {
	index := template.Must(template.ParseFiles("../html/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		UserInput := r.FormValue("UserInput")
		Crypto := UserInput
		url := "https://api.coingecko.com/api/v3/simple/price?ids=" + Crypto + "&vs_currencies=usd"
		url_2 := "https://api.coingecko.com/api/v3/coins/" + Crypto + "/market_chart?vs_currency=usd&days=1&interval=daily"
		url_3 := "https://api.coingecko.com/api/v3/search?query=" + Crypto
		url_4 := "https://api.coingecko.com/api/v3/search/trending"
		url_5 := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
		url_6 := "https://api.coingecko.com/api/v3/coins/" + Crypto
		body_price := traitement(url)
		body_interval := traitement(url_2)
		body_image := traitement(url_3)
		body_tendance := traitement(url_4)
		body_price_btc := traitement(url_5)
		coins := traitement(url_6)
		affichage_price := tracker("usd", body_price, false, true)
		affichage_interval := Simplify(Pourcentage(affichage(body_interval)))
		affichage_image := tracker("thumb", body_image, true, false)
		affichage_image2 := tracker("thumb", body_tendance, true, false)
		affichage_tendance := tracker("id", body_tendance, false, false)
		affichage_vote_posit := tracker("sentiment_votes_up_percentage", coins, false, false)
		affichage_vote_nega := tracker("sentiment_votes_down_percentage", coins, false, false)
		affichage_prix_tendance, err := strconv.ParseFloat(tracker("price_btc", body_tendance, false, false), 64)
		affichage_bitcoin, err := strconv.ParseFloat((tracker("usd", body_price_btc, false, true)), 64)
		affichage_prix_tendance_simp := (Price_Convert(affichage_bitcoin, false, affichage_prix_tendance))
		if err != nil {
			fmt.Println(err)
			return
		}
		data := crypto{
			Price:          affichage_price,
			Price_btc:      affichage_bitcoin,
			Interval:       affichage_interval,
			Image:          affichage_image,
			Tendance:       affichage_tendance,
			Image2:         affichage_image2,
			Price_tendance: affichage_prix_tendance_simp,
			Votes_P:        affichage_vote_posit,
			Votes_N:        affichage_vote_nega,
			Name:           strings.Title(UserInput),
		}

		index.Execute(w, data)
	})
	css := http.FileServer(http.Dir("../css/"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	http.ListenAndServe(":8180", nil)
}
