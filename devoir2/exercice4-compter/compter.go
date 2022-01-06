package compter
import "fmt"
/*
La fonction compter(n, m) doit retourner une chaîne de caractère qui liste tous
les entiers dans l'ordre croissant de n à m puis dans l'ordre décroissant de m
à n, séparés par des signes "+". (Si n > m, ceci n'est pas possible et la fonction
retournera donc une chaîne vide.)

Vous ne devez pas utiliser de boucles, la fonction compter sera donc récursive.

# Entrées
- n : un entier
- m : un entier

# Sortie
- s : une chaîne de caractères

# Exemples
compter(5, 5) = "5"
compter(2, 5) = "2+3+4+5+4+3+2"
*/

func compter(n, m uint) (s string) {
	if n > m {
		return ""
	}
	if n == m {
		return fmt.Sprint(m)
	}
	return croissant(n, m, n, "")
}

func croissant(n, m, og uint, s string) string {
	if n > m {
		return decroissant(n-2, og, s)
	}
	return croissant(n+1, m, og, s + fmt.Sprint(n) + "+")
}

func decroissant(n, m uint, s string) string {
	if n <= m {
		return s + fmt.Sprint(n)
	}
	return decroissant(n -1, m, s + fmt.Sprint(n) + "+")
}

/*
func aux(n, m , og uint, sens bool, s string) string {
	if sens {
		s += fmt.Sprint(n)

		if n == m {
			return aux(n-1, m, og, false, s)
		}
		if n == 0 {
			return aux(1, m, og, true, s + "+")
		}
		return aux(n+1, m, og, true, s + "+")
	}
	if n < og {
		return s
	}
	
	if n == 1 && og == 0{
		return s + "+" + fmt.Sprint(n) + "+0"
	}
	return aux(n-1, m, og, false, s + "+" + fmt.Sprint(n))
}
*/