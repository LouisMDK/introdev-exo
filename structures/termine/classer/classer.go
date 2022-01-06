package classer

/*
La fonction classer doit trier un tableau d'étudiants (tels que définis) ci-dessous
de celui qui a la meilleur moyenne (la plus élevée) à celui qui a la plus mauvaise
moyenne. En cas d'égalité de moyenne entre deux étudiants, celui qui a le nom de
famille qui arrive en premier dans l'ordre alphabétique doit être classer avant
l'autre (on peut utiliser <, >, <=, >=, == pour comparer les chaînes de caractères
par ordre alphabétique). Si deux étudiants ont la même moyenne et le même nom, on
met en premier celui qui a le prénom qui est en premier dans l'ordre alphabétique.

# Entrée
- promo : le tableau d'étudiants à trier

# Sortie
- classement : un tableau contenant les mêmes étudiants mais trié
*/
type etudiant struct {
	nom, prenom string
	moyenne     float64
}

func classer(promo []etudiant) (classement []etudiant) {
	if len(promo) == 0 || promo == nil {
		return classement
	}
	for i := 0; i < len(promo); i++ {
		classement = append(classement, promo[i])
	}
	for i := 0; i < len(promo); i++ {
		for j := 0; j < len(promo); j++ {
			if classement[i].moyenne > classement[j].moyenne {
				classement[i], classement[j] = classement[j], classement[i]
			}
			if classement[i].moyenne == classement[j].moyenne {
				if classement[i].nom < classement[j].nom {
					classement[i], classement[j] = classement[j], classement[i]
				}
				if classement[i].nom == classement[j].nom {
					if classement[i].prenom < classement[j].prenom {
						classement[i], classement[j] = classement[j], classement[i]
					}
				}
			}
		} 
	}
	return classement
}