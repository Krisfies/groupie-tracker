package main

import (
	"fmt"
	"html/template"
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
		body_price := Traitement(url)
		body_interval := Traitement(url_2)
		body_image := Traitement(url_3)
		body_tendance := Traitement(url_4)
		body_price_btc := Traitement(url_5)
		coins := Traitement(url_6)
		affichage_price := Tracker("usd", body_price, false, true)
		affichage_interval := Simplify(Pourcentage(Display(body_interval)))
		affichage_image := Tracker("thumb", body_image, true, false)
		affichage_image2 := Tracker("thumb", body_tendance, true, false)
		affichage_tendance := Tracker("id", body_tendance, false, false)
		affichage_vote_posit := Tracker("sentiment_votes_up_percentage", coins, false, false)
		affichage_vote_nega := Tracker("sentiment_votes_down_percentage", coins, false, false)
		affichage_prix_tendance, err := strconv.ParseFloat(Tracker("price_btc", body_tendance, false, false), 64)
		affichage_bitcoin, err := strconv.ParseFloat((Tracker("usd", body_price_btc, false, true)), 64)
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
	http.ListenAndServe(":8090", nil)
}
