package factorielle

/*
La fonction factorielle calcule la factorielle d'un nombre entier positif ou
nul. Pour rappel : 0! = 1 et pour n > 0, n! = n*(n-1)*(n-2)*...*1.

Vous ne devez pas utiliser de boucles, la fonction factorielle sera donc récursive

# Entrée
- n : l'entier dont on veut calculer la factorielle

# Sortie
- fact : la factorielle de n

# Exemple
factorielle(5) = 120
*/

func factorielle(n uint) (fact uint) {
	fact = aux(1, n+1, 1)
	return fact
}

func aux(start uint, end uint, n uint) (fact uint) {
	if start >= end {
		return n
	}
	return aux(start + 1, end, n*start)
}
