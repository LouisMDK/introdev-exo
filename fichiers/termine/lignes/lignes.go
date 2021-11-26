package lignes
import (
	"bufio"
	"os"
)
/*
La fonction lignes doit compter le nombre de lignes dans un fichier dont le nom
est indiqué en paramètre.

# Entrée
- fName : une chaîne de caractères correspondant à un nom de fichier

# Sortie
- nLignes : un entier indiquant le nombre de lignes dans le fichier de nom fName ou -1 si le fichier n'existe pas
*/

func lignes(fName string) (nLignes int) {

	var cFile *os.File
	var err error
	cFile, err = os.Open(fName)
	if err != nil {
		return -1
	}

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(cFile)

	for scanner.Scan() {
		nLignes++
	}


	return nLignes
}
