package monnaie

import (
	"errors"
)

var errPasAssezPaye error = errors.New("le montant payé est inférieur au montant de l'achat, impossible de rendre la monnaie")

/*
Étant donnés un montant d'achat en euros et un montant payé en euros,
la fonction rendreMonnaie retourne la somme qu'il faut rendre au client.

# Entrées
- eurosAchat : la partie entière du montant d'achat en euros
- centimesAchat : la partie décimale du montant d'achat (donc les centimes)
- eurosPayes : la partie entière du montant payé en euros
- centimesPayes : la partie décimale du montant payé (donc les centimes)

# Sorties
- eurosRendus : la partie entière du montant à rendre en euros
- centimesRendus : la partie décimale du montant à rendre (donc les centimes)
- err : une erreur qui doit être errPasAssezPaye si le montant payé est inférieur au montant de l'achat et nil sinon

# Exemple
rendreMonnaie(12, 20, 15, 0) = 2, 80
*/
func rendreMonnaie(eurosAchat, centimesAchat, eurosPayes, centimesPayes int) (eurosRendus, centimesRendus int, err error) {
	var centimes_achat int = eurosAchat*100 + centimesAchat
	var centimes_payes int = eurosPayes*100 + centimesPayes
	if centimes_payes < centimes_achat {
		return 0, 0, errPasAssezPaye
	}
	if centimes_payes == centimes_achat {
		return 0, 0, nil
	}
	var centimes_rendre int = centimes_achat - centimes_payes
	if centimes_rendre < 0 {
		centimes_rendre *= -1
	}
	eurosRendus = (centimes_rendre - centimes_rendre%100) / 100
	centimesRendus = centimes_rendre % 100
	return eurosRendus, centimesRendus, err
}
