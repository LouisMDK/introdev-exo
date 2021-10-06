package facteurspremiers
/*
La fonction facteursPremiers doit retourner un tableau contenant la liste de tous
les facteurs premiers d'un entier n, doublons compris

# Entrée
- n : un nombre entier positif

# Sortie
- facteurs : un tableau contenant tous les facteurs premiers de n, si n vaut 0 il
faut retourner un tableau à zéro éléments.

# Exemple
premiers(24) = [2 2 2 3] (l'ordre n'a pas d'importance)
*/
func facteursPremiers(n uint) (facteurs []uint) {

	if n < 0{
		panic("n doit être un nombre entier positif")
	}
	if n == 0{
		return []uint{}
	}
	if n == 1{
		return []uint{}
	}
	for{
		if n <= 1 {
			break
		}
		for i:=uint(2);i<=n;i++{
			if n % i == 0{
				facteurs = append(facteurs, i)
				n = n / i
				break
			}
		}

	}

	return facteurs
}
