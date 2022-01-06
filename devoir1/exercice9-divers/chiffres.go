package chiffres
import strconv
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
*/

func chiffres(n int) (c int) {
	return aux(0, n, -1)
}

func aux(start, end, n int) int {
	if start > end{
		var s string = strconv.Itoa(n)
		print(s)
		if len(s) > 10 {
			s = s[:10]
		}
		s = strconv.Atoi(s)
		return s
	}
	n += puissance(start)
	return aux(start+1, end, n)
}

func puissance(n int) (res int) {
	res = 1
	for i := 0; i < n; i++ {
		res *= n
	}
	return res
}