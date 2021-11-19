package puissant

/*
On dit qu'un mot est puissant s'il est la répétition d'un même motif 2 ou plus
de fois. Ainsi, "bonbon" est puissant (le même motif se répète 2 fois, on dit
aussi que c'est un carré) et "lalala" est puissant (le même motif se répète 3
fois, on dit aussi que c'est un cube). Par contre "bonjour" n'est pas puissant,
ni "bonbons" (le s final casse la répétition du motif "bon").

La fonction puissant doit déterminer si un mot, donné sous la forme d'une chaîne
de caractères ne comprenant que des lettres minuscules non accentuées, est
puissant ou non.

Des points seront donnés pour la réalisation d'une fonction aussi efficace ou
plus efficace que celle du prof.

# Entrée
- mot : le mot à analyser

# Sortie
- estPuissant : true si mot est puissant et false sinon

# Info
2021-2022, test 1, exercice 9
*/

func puissant(mot string) (estPuissant bool) {
	if len(mot) == 0 {
		return true
	}
	var motif []byte
	for i := 0; i < len(mot); i++{
		motif = append(motif, mot[i])
		if estMotif(motif, mot){
			return true
		}
	}
	return estPuissant
}


func estMotif(motif []byte, mot string) (bool) {
	var i int
	for {
		if i >= len(mot){
			break
		}
		for j := 0; j < len(motif); j++ {
			if i+j > len(mot) - 1{
				return false
			}
			if mot[i+j] != motif[j] {
				return false
			}
		}
		i += len(motif)
	}

	
	if len(motif) == len(mot){
		var same bool = true
		for i:=0;i<len(motif);i++{
			if motif[i] != mot[i] {
				same = false
				break
			}
		}
		if same{
			return false
		}
	}
	return true

}