package huitreines

/*
Le problème des huit reines consiste à placer, sur un échiquier (un tableau de 8 cases par 8 cases), 8 reines, de telle sorte qu'aucune d'entre-elles ne soit en position d'en manger une autre (c'est à dire de telle sorte qu'il n'y ait pas deux reines sur la même ligne, la même colonne ou la même diagonale de l'échiquier).

Ce problème se généralise en problème des n reines. On cherche alors à placer de la même façon n reines sur un échiquier de n cases par n cases.

Quand on analyse le problème, on se rend rapidement compte que chaque reine doit obligatoirement être sur une ligne différente des autres : il y a exactement une reine par ligne de l'échiquier. Partant de là, on peut procéder récursivement pour résoudre le problème :
1. placer une reine sur la première case de la première ligne (l) sans reine
2. si cette reine n'est pas mangée par une autre résoudre le même problème pour n-1 reines sur un échiquier ou les lignes 0 à l sont occupées
3. s'il y a une solution, c'est gagné
4. sinon (si la reine est mangée ou si le problème à n-1 reines n'a pas de solution), placer la reine sur la case suivante de la ligne (l) et recommencer à l'étape 2.

L'objectif de cet exercice est d'écrire une fonction huitreines qui résout le problème des n reines pour n donné.

# Entrée
- n : le nombre de reines à placer sur un échiquier de n cases par n cases

# Sortie
- plateau : un tableau de tableaux d'entiers de taille n*n qui contient des 0 là où il n'y a pas de reine et des 1 là où il y a des reines dans une solution du problème des n reines (si elle existe)
- ok : un booléen qui indique si le problème a une solution ou pas

# Exemple
huitreines(4) = [[0 1 0 0] [0 0 0 1] [1 0 0 0] [0 0 1 0]] (solution non unique)
*/
func huitreines(n int) (plateau [][]int, ok bool) {

	var plat [][]int = make([][]int, n, n)
	for i, _ := range plat{
		plat[i] = make([]int, n, n)
	}
	
	plateau = a(plat)

	return plateau, ok
}

func a(plateau [][]int) ([][]int){
	var l int = premiereLigneSansReine(plateau)
	for i:=0 ; i < len(plateau); i++ {
		plateau[l][i] = 1
		if estMange(plateau, l, i) {
			plateau[l][i] = 0
		}else{
			
		}
	}


	return plateau
}
func estMange(plateau [][]int, x, y int) (bool) {
	for i:=0; i<len(plateau); i++ {
		if plateau[x][i] == 1 && i != y{
			return false
		}
	}
	return true
}
func premiereLigneSansReine(plateau [][]int) (int){
	var isin bool
	for i, ligne := range plateau {
		isin = false
		for _, val := range ligne {
			if val == 1 {
				isin = true
				break
			}
			if !isin{
				return i
			}
		}
	}
	return -1
}