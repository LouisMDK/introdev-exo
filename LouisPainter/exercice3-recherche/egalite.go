package egalite

/*
On considère des ensembles de nombres représentés par des tableaux : on considère
que ces tableaux ne contiennent qu'une seule fois chaque nombre (puisqu'ils
représentent des ensembles) et les nombres ne sont pas nécessairement dans
l'ordre dans les tableaux.

On veut savoir si deux ensembles sont égaux ou pas, c'est-a-dire savoir si deux
tableaux contiennent les mêmes nombres ou pas. C'est à cette question que doit
répondre la fonction egalite.

# Entrées
- t1 : un tableau d'entiers (sans doublons) représentant un ensemble
- t2 : un tableau d'entiers (sans doublons) représentant un ensemble

# Sortie
- egaux : un booléen qui vaut true si t1 et t2 représentent le même ensemble et
          qui vaut false sinon
*/
func egalite(t1, t2 []int) (egaux bool) {
	var seen1 map[int]int = make(map[int]int) 
	var seen2 map[int]int = make(map[int]int)
	var val int
	var exists bool
	for i := 0; i < len(t1); i++ {
		val, _ = seen1[t1[i]]
		seen1[t1[i]] = val + 1
	}
	for j := 0; j < len(t2); j++ {
		val, _ = seen2[t2[j]]
		seen2[t2[j]] = val + 1
	}
	for i, k := range seen1 {
		val, exists = seen2[i]
		if !exists || val != k{
			return false
		}	
	}
	return true
}