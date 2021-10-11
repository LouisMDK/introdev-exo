package racaman
import "fmt"
/*
La suite de Racaman est définie par a(1) = 1, puis pour n > 1 par :
- a(n-1) - n si ce nombre est strictement supérieur à 0 et n'a encore jamais été
vu dans la suite
- a(n-1) + n sinon

La fonction racaman doit calculer a(n) pour n supérieur ou égal à 1.

# Entrée
- n : le numéro du terme de la suite à calculer

# Sortie
- an : la valeur du terme de la suite calculé (si ce terme n'est pas défini, on
retournera -1)

# Exemple
racaman(4) = 2
*/

var liste []int
func racaman(n int) (an int) {
	liste = nil
	if n <= 0 {
		return -1
	}
	val := sequence(n)
	fmt.Println(liste)
	return val
}

func sequence(n int) (an int) {
	if n == 1 {
		return 1
	}
	var val int
	val = sequence(n-1) - n
	if val > 0 && !contient(liste, val) {
		liste = append(liste, val)
		return val
	}
	val = sequence(n-1) + n
	return val
}
/*
func sequence(n int) (an int) {
	if n == 1 {
		return 1
	}
	var val int 
	val = sequence(n-1) - n
	if val > 0 && !dejaVu(val) {
		return val
	}
	return sequence(n-1) + n
}

func sequenceVals(n int, tab []int) (vus []int) {
	if n == 1 {
		tab = append(tab, 1)
		return tab
	}

	var val int
	var liste []int
	liste = sequenceVals(n-1, tab)
	val = liste[len(liste) - 1] - n
	if !contient(tab, val) && val > 0 {
		tab = append(tab, val)
		return tab
	}
	tab = append(tab, liste[len(liste)-1] + n)
	return tab
}
*/
func contient(liste []int, val int) (bool) {
	for _, valeur := range liste {
		if valeur == val {
			return true
		}
	}
	return false
}