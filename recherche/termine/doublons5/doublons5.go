package doublons5

/*
On dispose d'un tableau d'entiers et on suppose qu'il ne contient pas de
doublons (aucun nombre n'est présent plusieurs fois dans le tableau.)
On voudrait vérifier cela, ce que va faire la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui vaut true si le tableau ne contient pas de doublons et
false sinon
*/

func doublons(tab []int) (ok bool) {
	var vus map[int]int = make(map[int]int)

	for i := 0; i < len(tab); i++{
		b, _ := vus[tab[i]]
		vus[tab[i]] = b + 1
		if b + 1 > 1 {
			return false
		}
	}
	return true
}
