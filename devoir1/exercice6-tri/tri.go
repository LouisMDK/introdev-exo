package tri

/*
La fonction tri doit trier un tableau d'entiers du plus petit au plus grand.
Cette fonction ne doit pas modifier le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         petit au plus grand.
*/

func tri(tinit []int) (tfin []int) {
	tfin = make([]int, len(tinit))
	copy(tfin, tinit)
	TriRapide(tfin, 0, len(tfin))
	return tfin
}

func TriRapide(tab []int, p, r int) {
	if len(tab[p:r]) > 1 {
		var q = partitionner(tab, p, r)
		TriRapide(tab, 0, q)
		TriRapide(tab, q+1, r)
	}
}

func partitionner(tab []int, p, r int) int {
	var pivot int = tab[p]
	var i int = p
	var j int = r - 1

	for i < j {

		for i < j && tab[j] >= pivot {
			j--
		}

		if i < j {
			tab[i], tab[j] = tab[j], tab[i]
		}

		for i < j && tab[i] <= pivot {
			i++
		}

		if i < j {
			tab[i], tab[j] = tab[j], tab[i]
		}
	}
	return i
}