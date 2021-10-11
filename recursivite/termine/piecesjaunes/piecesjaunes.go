package piecesjaunes

import (
	"errors"
)

var errImpossible error = errors.New("impossible de constituer cette somme avec ces pièces")
var errListeVide error = errors.New("impossible de trouver le maximum dans une liste vide")
/*
Les pièces et billets que nous utilisons dans la vie de tous les jours ont des valeurs qui ont été choisies pour qu'il soit simple de rendre la monnaie (il suffit de toujours donner la pièce ou le billet de plus grande valeur possible, jusqu'à avoir rendu la somme voulue). Dans cet exercice, on imagine que ce n'est pas le cas : on dispose d'un ensemble de pièces de valeurs quelconques et on cherche à constituer une somme avec le moins de pièces possible.

La fonction meilleurDecomposition doit trouver la meilleur décomposition d'une somme d'argent (celle qui utilise le moins de pièces) en piochant des pièces dans l'ensemble donné en argument de la fonction.

# Entrées
- somme : la somme d'argent à constituer avec les pièces
- piecesDisponibles : un tableau des pièces disponibles

# Sorties
- pieces : un ensemble de pièces dont la somme vaut somme
- err : une erreur qui doit valoir errImpossible si la somme ne peut pas être réalisée à partir des pièces considérées.

# Exemples
meilleurDecomposition(15, []int{2, 3, 5, 5, 5}) = [5 5 5]
meilleurDecomposition(15, []int{2, 3, 5}) = errImpossible
*/

func meilleurDecomposition(somme int, piecesDisponibles []int) (pieces []int, err error) {
	if somme == 0 {
		return []int{}, nil
	}
	if somme < 0 {
		return pieces, errImpossible
	}
	var maxpos int = maxInferieurA(piecesDisponibles, somme)
	if maxpos == -1 {
		
		return pieces, errImpossible
	}
	pieces = append(pieces, piecesDisponibles[maxpos])
	somme -= piecesDisponibles[maxpos]
	piecesDisponibles = retirerElement(piecesDisponibles, maxpos)
	decomposition, er := meilleurDecomposition(somme, piecesDisponibles)
	if er != nil{
		return pieces, errImpossible
	}
	for _, val := range decomposition {
		if val != 0{
			pieces = append(pieces, val)
		}
	}

	return pieces, nil
}

func retirerElement(liste []int, pos int) (res []int){
	for i := 0; i<len(liste); i++{
		if i != pos{
			res = append(res, liste[i])
		}
	}
	return res
}
func maxInferieurA(liste []int, n int) (pos int){

	if len(liste) == 0 || liste == nil {
		return -1
	}
	pos = 0
	for i := 0; i<len(liste); i++ {
		if liste[i] > liste[pos] && liste[i] <= n{
			pos = i
		}
	}

	if pos == 0 && liste[pos] < n {
		return -1
	}
	return pos

}