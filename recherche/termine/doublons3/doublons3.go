package doublons3

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement k fois chaque nombre compris entre 1 et n/k. On voudrait vérifier
cela. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers
- k : un entier

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement k fois chaque
entiers de 1 à len(tab)/k et false sinon
*/

func doublons(tab []int, k int) (ok bool) {
	if len(tab) == 0 || tab == nil {
		return true
	}
	var n int = len(tab) / k
	var vus []int = make([]int, n)
	for i := 0; i < len(tab); i++ {
		if tab[i] > n {
			return false
		}
		vus[tab[i] - 1] += 1

	}
	for i := 0; i < len(vus); i++ {
		if vus[i] != k {
			return false
		}
	}
	return true
}
