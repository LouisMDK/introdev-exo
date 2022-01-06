package compte

/*
La fonction compte doit indiquer combien de fois une valeur v apparaît dans un
tableau tab.

Vous ne devez pas utiliser de boucles, la fonction compte sera donc récursive.

Vous pouvez vous baser sur la remarque suivante : le nombre de fois que v apparaît
dans tab est la somme du nombre de fois où v apparaît dans la première moitié
de tab et du nombre de fois où v apparaît dans la deuxième moité de tab.

# Entrées
- tab : un tableau d'entiers
- v : la valeur à chercher

# Sortie
- num : le nombre de fois que v apparaît dans tab
*/

func compte(tab []int, v int) (num int) {
	if len(tab) == 0 || tab == nil {
		return 0
	}
	if len(tab) == 1 {
		if tab[0] == v {
			return 1
		}
		return 0
	}

	var q int = len(tab) / 2
	var gauche int = compte(tab[0:q], v)
	var droite int = compte(tab[q:len(tab)], v)

	return gauche + droite
}