package racaman
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

var sequence []int
func racaman(n int) (an int) {
	sequence = []int{0}
	if n < 1 {
		return -1
	}
	return aux(1, n)
}

func aux(start, end int) (val int){
	val = sequence[start - 1] - start
	if val > 0 && !contient(val) {
		sequence = append(sequence, val)
	}else{
		val = sequence[start - 1] + start
		sequence = append(sequence, val)
	}
	

	if start == end {
		return val
	}
	return aux(start+1, end)
}

func contient(n int) bool {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] == n {
			return true
		}
	}
	return false
}
