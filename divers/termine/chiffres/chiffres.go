package chiffres

/*
1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317
Les 10 derniers chiffres de ce nombre sont 0405071317, qu'on représentera par
l'entier 405071317
La fonction chiffres doit donner l'entier qui représente, sur le même modèle,
les 10 derniers chiffres de 1^1 + 2^2 + ... + n^n pour un entier n donné en
argument.

Une fonction qui calcule chiffres(10000) en moins d'une minute sur ma machine
rapportera plus de points

# Entrée
- n : l'entier maximum utilisé dans le calcul

# Sorties
- c : l'entier représentant les 10 derniers chiffres de 1^1 + 2^2 + ... + n^n

Inspiré du problème 48 de projecteuler.net

# Info
2021-2022, test2, exercice 9
*/

func chiffres(n int) (c int) {
	for i := 1; i <= n; i++ {
		c += puissance(i)
		c %= 10000000000
	}
	return c
}

func puissance(n int) int {
	var res int = 1
	for i := 1; i <= n; i++ {
		res *= n 
		res %= 10000000000
	}
	return res
}