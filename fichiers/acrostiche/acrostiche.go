package acrostiche
import (
        "os"
        "bufio"
        "fmt"
        "log"
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
        // Ouverture du fichier
	var myFile *os.File
	var err error
	myFile, err = os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}

	// Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)
	// Lecture des lignes du fichier
        var lettre string
	for scanner.Scan() {
                lettre = scanner.Text()
                if lettre != " " {
                        mot = mot + lettre[0].String()   
                }
                
	}

	// Vérification que tout s'est bien passé
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	return mot
}
