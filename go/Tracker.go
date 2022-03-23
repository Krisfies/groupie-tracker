package main

// Le principal outil que nous avons crées afin de trier les structures de l'api
func Tracker(arg string, body string, Is_image bool, Is_number bool) string {
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
