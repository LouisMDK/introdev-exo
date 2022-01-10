package nombre

import (
    "os"
    "bufio"
)

/*
La fonction nombre doit dire si le fichier dont le chemin est donné en paramètre
contient ou ne contient pas le chiffre 1. On suppose que le fichier en question
existe toujours (vous n'avez pas à vous préocupper des erreurs à l'ouverture du
fichier).

# Entrée
- chemin : le chemin pour accéder au fichier à traiter

# Sortie
- contient : un booléen qui vaut true si le chiffre 1 apparaît dans le fichier
             considéré et false sinon (le nom du fichier ne compte pas, seul
             son contenu est à prendre en compte).
*/


func nombre(chemin string) (contient bool) {
    var cFile *os.File
    var err error

    cFile, err = os.Open(chemin)

    if err != nil {
        return false
    }

    var scanner *bufio.Scanner = bufio.NewScanner(cFile)
    var ligne string
    
    for scanner.Scan() {
        ligne = scanner.Text()
        for i := 0; i < len(ligne); i++ {
            if string(ligne[i]) == "1" {
                return true
            }
        }
    }
    err = cFile.Close()
	return false
}