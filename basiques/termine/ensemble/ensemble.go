package ensemble

/*
Un ensemble d'entier est un paquet de plusieurs entiers, sans doublons.
La fonction estEnsemble doit indiquer si en tableau d'entiers correspond à un
ensemble ou non.

# Entrée
- E : un tableau d'entiers

# Sortie
- b : un booléen indiquant si E est bien un ensemble ou non (nil n'est pas un ensemble)

# Exemples
estEnsemble([]int{1, 2, 3}) = true
estEnsemble([]int{1, 2, 2}) = false
*/
func estEnsemble(E []int) (bool) {
	if E == nil{
		return false
	}
	for i:=0;i<len(E);i++{
		for j:=0;j<len(E);j++{
			if E[i] == E[j] && i != j{
				return false
			}
		}
	}
	return true
}
