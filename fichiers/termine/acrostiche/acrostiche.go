package acrostiche
import (
	"os"
	"bufio"
)
/*
La fonction acrostiche doit reconstituer le mot formé par le premier caractère
de chaque ligne d'un fichier, en ignorant les lignes vides.

# Entrée
- fName : le nom d'un fichier

# Sortie
- mot : la chaîne de caractère obtenue en mettant l'une après l'autre, dans
        l'ordre des lignes du fichier, les premières lettres de chacunes des
        lignes du fichier dont le nom est fName (on considère que ce fichier
        existe toujours)
*/

func acrostiche(fName string) (mot string) {
	var cFile *os.File
	var err error

	cFile, err = os.Open(fName)
	if err != nil {
		return ""
	}
	var scanner *bufio.Scanner = bufio.NewScanner(cFile)
	var ligne string
	for scanner.Scan() {
		ligne = scanner.Text()
		if len(ligne) > 0 {
			mot += string(ligne[0])
		}
	}
	return mot
}
