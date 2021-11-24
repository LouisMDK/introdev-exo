package nombrespremiers
/*
La fonction premiers doit retourner un tableau contenant les nombres
premiers compris entre 0 et un entier n.

# Entrée
- n : un nombre entier

# Sortie
- p : un tableau contenant tous les nombres premiers compris entre 0 et n inclus, s'il n'existe pas de tels nombres premiers, il faut que p soit un tableau contenant 0 éléments

# Exemple
premiers(10) = [2 3 5 7] (l'ordre n'a pas d'importance)
*/
func premiers(n int) (p []int) {
	var premier bool
	for i:=2;i<=n;i++ {
		premier = true
		for j:=2;j<i;j++{
			if i % j == 0{
				premier = false
			}
		}
		if premier{
			p = append(p, i)
		}
	}
	if len(p) == 0{
		return make([]int, 0)
	}
	return p
}
