package monnaie2

import (
	"errors"
)

var errPasAssezPaye error = errors.New("le montant payé est inférieur au montant de l'achat, impossible de rendre la monnaie")

/*
Étant donnés un montant d'achat en euros et un montant payé en euros,
la fonction rendreMonnaie retourne une liste de pièces et billets qui
permet de rendre au client la somme qu'il faut.

# Entrées
- eurosAchat : la partie entière du montant d'achat en euros
- centimesAchat : la partie décimale du montant d'achat (donc les centimes)
- eurosPayes : la partie entière du montant payé en euros
- centimesPayes : la partie décimale du montant payé (donc les centimes)

# Sorties
- eurosRendus : un tableau contenant les valeurs des pièces et billets en euros à rendre, ces valeurs doivent être prises parmi 1, 2, 5, 10, 20, 50, 100, 200 et 500
- centimesRendus : un tableau contenant les valeurs des pièces en centimes à rendre, ces valeurs doivent être prises parmi 1, 2, 5, 10, 20 et 50
- err : une erreur qui doit être errPasAssezPaye si le montant payé est inférieur au montant de l'achat et nil sinon

# Exemple
rendreMonnaie(12, 20, 15, 0) = [1, 1], [50, 20, 10] (ce n'est pas la seule possibilité pour ce rendu)
*/
func rendreMonnaie(eurosAchat, centimesAchat, eurosPayes, centimesPayes int) (eurosRendus, centimesRendus []int, err error) {
	var centimes_achat int = eurosAchat*100 + centimesAchat
	var centimes_payes int = eurosPayes*100 + centimesPayes
	if centimes_payes < centimes_achat {
		return []int{}, []int{}, errPasAssezPaye
	}
	if centimes_payes == centimes_achat {
		return []int{}, []int{}, nil
	}
	var centimes_rendre int = centimes_achat - centimes_payes
	if centimes_rendre < 0 {
		centimes_rendre *= -1
	}
	var max_a_rendre int
	
	for _, val := range []int{500, 200, 100, 50, 20, 10, 5, 2, 1}{
		max_a_rendre = (centimes_rendre - (centimes_rendre % (val * 100))) / 100
		if max_a_rendre > 0{
			eurosRendus = append(eurosRendus, max_a_rendre)
			centimes_rendre -= max_a_rendre * 100
		}
	}
	for _, val := range []int{50, 20, 10, 5, 2, 1}{
		max_a_rendre = centimes_rendre - centimes_rendre % val
		if max_a_rendre > 0{
			centimesRendus = append(centimesRendus, max_a_rendre)
			centimes_rendre -= max_a_rendre
		}
	}
	
	return eurosRendus, centimesRendus, err
}
