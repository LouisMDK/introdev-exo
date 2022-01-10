package souschaine
/*
Une chaine sc est une sous chaine d'une chaine s si on peut reconstruire s en rajoutant des lettres à sc. Une chaine sc est une sous chaine commune à s1 et s2 si c'est une sous chaine de s1 et une sous chaine de s2. C'est une plus longue sous chaine commune si on ne peut pas trouver sc2 qui soit aussi une sous chaine commune à s1 et s2 et qui ait strictement plus de lettres que sc.

Un algorithme a été vu en cours pour calculer la longueur de plus longues sous chaines communes. L'objet de cet exercice est de l'adapter pour obtenir, en plus de leur longueur, les sous chaines elles-mêmes.

# Entrées
- s1 : une chaîne
- s2 : une chaîne

# Sortie
- s : une plus longue sous chaine commune à s1 et s2

# Exemple
sousChaine("bonjour", "bonsoir") = bonor
*/
var souschaines [][]int

func sousChaine(s1, s2 string) (res string) {
	if s1 == s2 {
		return s1
	}
	if s1 == "" || s2 == "" {
		return ""
	}

	souschaines = make([][]int, len(s1) + 1)
	for i:=0; i<len(souschaines); i++{
		souschaines[i] = make([]int, len(s2) + 1)

		for j := 0; j < len(souschaines[i]); j++ {
			if i == 0 {
				souschaines[i][j] = 0
			}else if j == 0 {
				souschaines[i][j] = 0
			}else {
				souschaines[i][j] = -1
			}
		} 
	}
	plusLongueSousChaine(s1, s2)
	res = ""
	var start int = 0
	for i := 0; i < len(souschaines); i++ {
		for j := 0; j < len(souschaines[i]); j++ {
			if souschaines[i][j] == start {
				if len(s1) > len(s2) {
					res += string(s2[j])
				} else {
					res += string(s1[i])
				}
				start++
				break
			}
		}
	}
	return res[:len(res)-1]
}

func plusLongueSousChaine(s1, s2 string) (int){

	if souschaines[len(s1)][len(s2)] < 0 {
		if s1[len(s1) - 1] == s2[len(s2) - 1]{
			souschaines[len(s1)][len(s2)] = 1 + plusLongueSousChaine(s1[:len(s1) - 1], s2[:len(s2) - 1])

		}else{
			souschaines[len(s1)][len(s2)] = max(plusLongueSousChaine(s1[:len(s1)], s2[:len(s2) - 1]), plusLongueSousChaine(s1[:len(s1) - 1], s2[:len(s2)]))
		}
	}
	
	return souschaines[len(s1)][len(s2)]

}

func max(n, m int) (int){
	if n > m {
		return n
	}
	return m
}