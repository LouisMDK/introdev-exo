package alphabetique
/*
La fonction alphabetique trie un tableau de chaînes de caractères dans l'ordre
alphabétique.

# Entrée
- dico : le tableau de chaînes de caractères à trier
*/

func alphabetique(dico []string) {
	for i := 0; i < len(dico); i++ {
		for j := 0; j < len(dico); j++ {
			if alphabet(dico[i], dico[j]) {
				dico[i], dico[j] = dico[j], dico[i]
			}
		}
	}
}

func alphabet(m1, m2 string) (m1avantm2 bool) {
	if m2 == "" {
		return false
	}
	if m1 == "" {
		return true
	}
	var length = min(len(m1), len(m2))
	for z := 0; z < length; z++ {

		if m1[z] > m2[z] {
			return false
		}
		if m1[z] < m2[z] {
			return true
		}
	}
	return len(m1) < len(m2)
}

func min(a, b int) (int) {
	if a > b {
		return b
	}
	return a
}