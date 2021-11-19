package doublons4
/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement k fois chaque nombre compris entre 1 et n/k. On voudrait vérifier
cela, mais on ne connaît pas k. La fonction doublons va calculer k, si c'est
possible.

# Entrée
- tab : un tableau d'entiers

# Sortie
- k : un entier (s'il existe) tel que tab contient exactement k fois chaque
entiers de 1 à len(tab)/k
- ok : un booléen qui indique si k existe ou non
*/

func doublons(tab []int) (k int, ok bool) {
	if len(tab) == 0 || tab == nil {
		return 0, true
	}
	var vus []int = make([]int, len(tab))

	for i := 0; i < len(tab); i++ {
		if tab[i] > len(tab) || tab[i] < 1 {
			return 0, false
		} 
		vus[tab[i] - 1] += 1
	}
	k = vus[0]
	var warning bool
	for i := 0; i < len(vus); i++ {
		if vus[i] == 0 {
			if warning {
				break
			}
			warning = true
		}else {
			if vus[i] != k || warning {
				return 0, false
			}
		}
	}
	return k, true

}
