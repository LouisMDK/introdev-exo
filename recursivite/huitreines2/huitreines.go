package huitreines2
import "fmt"
/*
Le problème des huit reines consiste à placer, sur un échiquier (un tableau de 8 cases par 8 cases), 8 reines, de telle sorte qu'aucune d'entre-elles ne soit en position d'en manger une autre (c'est à dire de telle sorte qu'il n'y ait pas deux reines sur la même ligne, la même colonne ou la même diagonale de l'échiquier).

Ce problème se généralise en problème des n reines. On cherche alors à placer de la même façon n reines sur un échiquier de n cases par n cases.

Quand on analyse le problème, on se rend rapidement compte que chaque reine doit obligatoirement être sur une ligne différente des autres : il y a exactement une reine par ligne de l'échiquier. Partant de là, on peut procéder récursivement pour résoudre le problème :
1. placer une reine sur la première case de la première ligne (l) sans reine
2. si cette reine n'est pas mangée par une autre résoudre le même problème pour n-1 reines sur un échiquier ou les lignes 0 à l sont occupées
3. s'il y a une solution, c'est gagné
4. sinon (si la reine est mangée ou si le problème à n-1 reines n'a pas de solution), placer la reine sur la case suivante de la ligne (l) et recommencer à l'étape 2.

L'objectif de cet exercice est d'écrire une fonction huitreines qui trouve toutes les solutions au problème des n reines pour n donné.

# Entrée
- n : le nombre de reines à placer sur un échiquier de n cases par n cases

# Sortie
- plateaux : un tableau de tableaux de tableaux d'entiers, chacun de taille n*n et contenant des 0 là où il n'y a pas de reine et des 1 là où il y a des reines dans une solution du problème des n reines (si elle existe), sans doublons

# Exemple
huitreines(1) = [[[1]]]
huitreines(2) = []
*/

var g [][][]int
func huitreines(n int) (plateaux [][][]int) {
	g = [][][]int{}
	var plateau [][]int
	plateau = make([][]int, n)
	for i := 0; i < len(plateau); i++ {
		plateau[i] = make([]int, n)
	}
	aux(plateau)
	fmt.Println(g)
	return g
}

func aux(plat [][]int) (plt [][]int, ok bool) {
	var l int = premiereLigne(plat)
	if l == -1 {
		g = append(g, plat)
		return plat, true
	}
	var next [][]int = make([][]int, len(plat))
	var tmp [][]int = make([][]int, len(plat))

	for i := 0; i < len(plat[l]); i++ {
		plat[l][i] = 1
		if estMange(plat, l, i) {
			plat[l][i] = 0
		} else {
			copy(next, plat)
			tmp, ok = aux(next)
			if ok {
				copy(next, tmp)
			} else {
				plat[l][i] = 0
			}
		}
	}
	return next, ok
}

func estMange(plateau [][]int, l, c int) bool {
	for ll := 0; ll < len(plateau); ll++ {
		if ll != l {
			if plateau[ll][c] == 1 {
				return true
			}
			for cc := 0; cc < len(plateau[ll]); cc++ {
				if plateau[ll][cc] == 1 {
					diffL := ll - l
					diffC := cc - c
					if diffL == diffC || diffL == -diffC {
						return true
					}
				}
			}
		}
	}
	return false
}

func premiereLigne(plat [][]int) int {
	var clear bool
	for i := 0; i < len(plat); i++ {
		clear = true
		for j := 0; j < len(plat[i]); j++ {
			if plat[i][j] == 1 {
				clear = false
			}
		}
		if clear {
			return i
		}
	}
	return -1
}
