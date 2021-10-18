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
	if n < 1 {
		return -1
	}
	return sequence(n)
}

func sequence(n int) (int) {
	fmt.Println(liste, n)
	if n == 1 {
		return 1
	}
	var val int = sequence(n - 1) - n
	if val > 0 && !contient(liste, val) {
		liste = append(liste, val)
		return val
	}
	val = sequence(n - 1) + n
	return val
}

func contient(l []int, v int) (bool) {
	for i := 0; i < len(l); i++ {
		if l[i] == v {
			return true
		}
	}
	return false
}