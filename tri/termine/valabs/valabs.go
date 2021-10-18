package valabs

/*
La fonction valabs doit trier un tableau d'entiers de la plus petite valeur absolue
à la plus grande valeur absolue. En cas
d'égalité de valeur absolue, les nombres
négatifs doivent être placés avant les
nombres positifs.

# Entrée
- tab : un tableau d'entiers
*/

func valabs(tab []int) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab); j++ {
			
			var v1 int = tab[i]
			if v1 < 0 {
				v1 *= -1
			}

			var v2 int = tab[j]
			if v2 < 0 {
				v2 *= -1
			}

			if v1 != v2 {
				if v1 < v2 {
					tab[i], tab[j] = tab[j], tab[i]
				}
			}else{
				if tab[i] < tab[j] {
					tab[i], tab[j] = tab[j], tab[i]
				}
			}
		}
	}
}
