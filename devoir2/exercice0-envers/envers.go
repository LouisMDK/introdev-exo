package envers

/*
Étant donné un tableau, la fonction envers doit l'inverser (en place), c'est à
dire que le premier élément du tableau doit devenir le dernier, le deuxième
élément doit devenir l'avant dernier, etc jusqu'au dernier élément qui doit
devenir le premier.

La fonction modifie le tableau d'entrée et n'a donc pas de sorties.

# Entrée
- tab : le tableau à inverser

# Exemple
envers appelée sur le tableau [1 2 3] doit le transformer en [3 2 1]
*/

func envers(t []int) {
	for i := 0; i < len(t) / 2; i++ {
		t[i], t[len(t) - i - 1] = t[len(t) - i - 1], t[i]
	}
}

/*
func envers2(t []int) {
	var inverse []int = make([]int, len(t))
	for i := 0; i < len(t); i++ {
		inverse[i] = t[len(t) - i - 1]
	}
	copy(t, inverse)
}
*/