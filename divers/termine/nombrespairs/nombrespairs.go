package nombrespairs

/*
La fonction sommeNombresPairs doit retourner la somme de tous les nombres
pairs contenus entre deux entiers (inclus).

# Entrées
- a : un nombre entier, une des bornes
- b : un nombre entier, l'autre borne

# Sortie
- somme : la somme de tous les nombres entiers compris entre a et b inclus

# Exemple
sommeNombresPairs(2, 7) = 12
*/
func sommeNombresPairs(a, b int) (somme int) {
	if a > b{
		b, a = a, b
	}
	for i:=a;i<=b;i++{
		if i % 2 == 0{
			somme += i
		}
	}
	return somme
}
