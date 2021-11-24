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

# Info
2021-2022, test2, exercice 3
*/

func egalite(t1, t2 []int) (egaux bool) {
	var seen1 map[int]int = make(map[int]int)
	var seen2 map[int]int = make(map[int]int)

	if len(t1) != len(t2){
		return false
	}
	var b int
	for i := 0; i < len(t1); i++ {
		b, _ = seen1[t1[i]]
		seen1[t1[i]] = b + 1
		b, _ = seen2[t2[i]]
		seen2[t2[i]] = b + 1
	}
	for key, val := range seen1 {
		b, _ = seen2[key]
		if b != val {
			return false
		}
	}
	return true
}
