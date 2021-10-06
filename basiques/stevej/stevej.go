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

	if a < 1900 || a > 2021{
		return 0, errPasUneDate
	}
	if m < 1 || m > 12{
		return 0, errPasUneDate
	}
	if j < 1{
		return 0, errPasUneDate
	}
	
	if (a % 4 == 0 && a % 100 != 0) || a % 400 == 0{ // l'année est bissextile
		if m == 2 && j > 29{
			return 0, errPasUneDate
		}
	}else{
		if m == 2 && j > 28{
			return 0, errPasUneDate
		}
	}
	if m == 4 || m == 6 || m == 9 || m == 11{
		if j > 30{
			return 0, errPasUneDate
		}
	}else{
		if j > 31{
			return 0, errPasUneDate
		}
	}
	if a == 1955 && m == 2 && j == 24{
		return 0, nil
	}
	
	var before bool = false
	if a < 1955 {
		before = true
	}
	if a == 1995 && m < 2{
		before = true
	}
	if a == 1995 && m == 2 && j < 24{
		before = true
	}
	
	if before{

		


	}else{
		if a == 1955{
			if m == 2{
				return j - 24, nil
			}
			ej += 5+j
			for i:=uint(2);i<m;i++{
				if i == 4 || i == 6 || i == 9 || i == 11{
					ej+=30
				}else{
					ej+=31
				}
			}
		}

	}
	return ej, err
}
