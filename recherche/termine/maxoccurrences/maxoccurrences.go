package maxoccurrences

/*
Étant donné un tableau d'entiers t, la fonction maxoccurrences doit retourner
l'entier qui apparaît le plus souvent dans t et le nombre de fois que cet
entier apparaît. En cas d'égalité on choisira arbitrairement l'un des entiers
à égalité. Si le tableau est vide on retournera un entier quelconque et 0 pour son nombre d'apparitions.

# Entrées
- t : le tableau dans lequel chercher

# Sortie
- n : l'entier qui apparaît le plus de fois dans t
- occ : le nombre de fois que n apparaît dans t

# Exemple
maxoccurrences([]int{1, 2, 3, 4, 3}) = 3, 2
*/
func maxoccurrences(t []int) (n int, occ int) {

	if t == nil || len(t) == 0 {
		return -1, 0
	}
	for _, val := range t {
		if val > n {
			n = val
		}
	}
	var occurences []int = make([]int, n+1)
	n = 0
	for _, val := range t {
		occurences[val] = occurences[val] + 1
		if occurences[val] > occ {
			occ = occurences[val]
			n = val
		}
	}
	return n, occ
}
