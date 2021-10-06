package nombrespremiers2

/*
La fonction selectionPremiers filtre le contenu d'un tableau d'entiers pour ne garder que ceux qui sont des nombres premiers, en éliminant les doublons.

# Entrée
- t : un tableau d'entiers

# Sortie
- p : un tableau contenant tous les nombres premiers de t, sans doublons.
Si t est vide, p doit être identique à t.

# Exemple
selectionPremiers([]int{1, 2, 2, 3, 4, 5}) = [2 3 5] (l'ordre n'a pas d'importance)
*/

func selectionPremiers(t []int) (p []int) {
	if t == nil || len(t) == 0{
		return t
	}
	for _, val := range(t){
		if val > 0 && estPremier(val){

			if !contient(p, val){
				p = append(p, val)
			}
			
		}
	}
	if len(p) == 0{
		return []int{}
	}
	return p
}

func contient(p []int, n int) (bool){
	var res bool
	for _, val := range(p){
		if val == n{
			return true
		}
	}
	return res
}

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