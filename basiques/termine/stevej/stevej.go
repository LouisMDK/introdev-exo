package stevej

import (
	"errors"
)

var errPasUneDate error = errors.New("la date indiquée n'est pas valide")

/*
Steve Jobs est né le 24 février 1955. La fonction ecart devra indiquer le nombre
de jours de différence entre une date (à partir du 1er janvier 1900 jusqu'au 31
décembre 2021) et la date de naissance de Steve Jobs. L'écart en jours sera toujours
un nombre positif ou nul (on ne se préoccupera pas de savoir si la date est antérieure
ou postérieure à la naissance de steve Jobs).

# Entrées
- j : un numéro de jour
- m : un numéro de mois
- a : un numéro d'année

# Sorties
- ej : l'écart en jours entre j/m/a et le 24/2/1955
- err : une erreur qui vaudra errPasUneDate quand j/m/a n'est pas une date valide et nil sinon

# Exemple
ecart(17, 4, 1986) = 11375
*/
func ecart(j, m, a uint) (ej uint, err error) {

	var annee []uint = []uint{31,28,31,30,31,30,31,31,30,31,30,31}
	var anneeb []uint = []uint{31,29,31,30,31,30,31,31,30,31,30,31}
	if m > 12 {
		return ej, errPasUneDate
	}

	if a % 400 == 0 || (a % 4 == 0 && a % 100 != 0) {
		if a < 1900 || a > 2021 || j == 0 || m == 0 || j > anneeb[m-1] {
			return ej, errPasUneDate
		}
		for i := uint(0); i<m-1; i++ {
			ej += anneeb[i]
		}
	}else{
		if a < 1900 || a > 2021 || j == 0 || m == 0 || j > annee[m-1] {
			return ej, errPasUneDate
		}
		for i := uint(0); i < m-1; i++ {
			ej += annee[i]
		}
	}

	ej += j

	for i := uint(1900); i < a; i++ {
		if (i % 4 == 0 && i % 100 != 0) || (i % 400 == 0){
			ej += 366
		}else{
			ej += 365
		}
	}

	if ej >= 20143 {
		return ej - 20143, err
	}
	return 20143 - ej, err
}
