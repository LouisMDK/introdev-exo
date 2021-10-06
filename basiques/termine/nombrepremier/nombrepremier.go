package nombrepremier

/*
La fonction estPremier doit indiquer  par un booléen si un nombre est premier
ou pas

# Entrée
- n : un nombre entier

# Sortie
- b : un booléen indiquant si n est premier ou pas

# Exemples
estPremier(5) = true
estPremier(10) = false
*/
func estPremier(n int) (b bool) {
	if n == 0 || n == 1{
		return false
	}
	for i:=2;i<n;i++ {
		if n % i == 0{
			return false
		}
	}
	return true
}
