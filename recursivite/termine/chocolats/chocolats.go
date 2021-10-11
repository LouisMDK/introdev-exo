package chocolats
/*
Une marque de barres de chocolat fait une promotion~: si on retourne k embalages on en obtient une gratuite. On se demande alors combien de barres de chocolat on peut obtenir quand on dispose de n euros et que chacune coûte m euros.
La fonction miam doit répondre (de manière récursive) à cette question.

# Entrées
- n : la somme dont on dispose en euros
- m : le prix d'une barre de chocolat en euros, toujours supérieur à 0
- k : le nombre d'embalages qu'il faut pour obtenir une barre de chocolat gratuite, toujours supérieur à 0

# Sortie
- choco : le nombre de barre de chocolat qu'on peut obtenir en tout

# Exemple
*/
func miam(n, m, k uint) (choco uint) {

	var barres_achetes uint = chocobarres(n, m)
	var barres_bonus uint = chocobonus(barres_achetes, k)
	return barres_achetes + barres_bonus
}

func chocobarres(n, m uint) (choco uint) {
	if n < m {
		return 0
	}
	return 1 + chocobarres(n-m, m)
}

func chocobonus(barres, k uint) (choco uint) {
	if barres < k {
		return 0
	}
	var reste uint = barres % k
	var papier uint = (barres - reste) / k
	return papier + chocobonus(papier + reste, k)
}