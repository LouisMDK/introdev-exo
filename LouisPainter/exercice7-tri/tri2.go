package tri2
/*
La fonction triabs doit trier un tableau d'entiers de la plus grande valeure
absolue à la plus petite valeure absolue. Cette fonction ne doit pas modifier
le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         grand (en valeure absolue) au plus petit (en valeure absolue).
*/

func triabs(tinit []int) (tfin []int) {
	for i := 0; i < len(tinit); i++ {
		tfin = append(tfin, tinit[i])
	}
	triRapide(tfin, 0, len(tfin))
	return tfin
}


func triRapide(tab []int, p, r int) {
	if len(tab[p:r]) > 1 {
		var q = partitionner(tab, p, r)
		triRapide(tab, 0, q)
		triRapide(tab, q+1, r)
	}
}

func partitionner(tab []int, p, r int) int {
	var pivot = absolue(tab[p])
	var i = p
	var j = r - 1

	for i < j {

		for i < j && absolue(tab[j]) <= pivot {
			j--
		}
		if i < j {
			tab[i], tab[j] = tab[j], tab[i]
		}
		for i < j && absolue(tab[i]) >= pivot {
			i++
		}
		if i < j {
			tab[i], tab[j] = tab[j], tab[i]
		}
	}
	return i
}

func absolue(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}