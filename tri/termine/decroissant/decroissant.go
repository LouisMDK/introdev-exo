package decroissant

/*
La fonction décroissant doit trier un tableau d'entiers du plus grand
au plus petit.

# Entrée
- tab : le tableau à trier
*/

func decroissant(tab []int) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab); j++ {
			if tab[j] < tab[i] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
}
