package prefixes

/*
La fonction numPrefixes doit compter le nombre de chaînes de caractères dans un
tableau qui commencent par une chaîne donnée.

# Entrées
- t : un tableau de chaînes de caractères
- s : une chaîne de caractères

# Sorties
- n : le nombre de chaînes de t qui commencent par s

# Exemple
numPrefixes([]string{"bonjour", "bonsoir", "salut", "bye bye"}, "bon") = 2
*/

func estPrefixe(c string, s string) (b bool){
	taille_s := len(s)
	taille_c := len(c)
	for i:=0;i<taille_s;i++{
		if i > taille_c - 1{
			return false
		}
		if c[i] != s[i]{
			return false
		}
		
	}
	return true
}

func numPrefixes(t []string, s string) (n int) {
	if len(t) == 0{
		return 0
	}
	if len(s) == 0{
		return len(t)
	}
	for _, chaine := range t{
		if estPrefixe(chaine, s){
			n++
		}
	}
	return n
}
