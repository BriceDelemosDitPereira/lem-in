package anthil

import (
	"fmt"
)

// Permet de déplacer la fourmi dans la salle suivante
func (ant *Ant) Next_Room(tab_path_used *[][2]*Room) string {
	// tab_path_used est un pointeur vers un tableau de tableau où chaque tableaux internes est un tableau de 2 pointeurs vers des objets Room
	// En gros c'est une liste de path où chaque path est représenté par 2 Room
	if ant.Is_End {
		return ""
	}
	curentpath := [2]*Room{ant.Path[ant.Pointeur], ant.Path[ant.Pointeur+1]}
	if !Path_Not_Used(tab_path_used, curentpath) || ant.Path[ant.Pointeur+1].Ifant { // si curentpath est déjà dans tab_path_used ou si il y a déjà une fourmi dans la Room suivante
		return ""
	}
	ant.Path[ant.Pointeur].Ifant = false         // pas de fourmis dans la Room actuelle car on va la déplacer dans la case suivante
	ant.Pointeur++                               // on déplace la fourmi dans la salle suivante, c'est à dire de la salle ant.Path[Pointeur] à la salle ant.Path[Pointer+1]
	if ant.Path[ant.Pointeur].Type_Room == End { // si la room est la dernière room
		ant.Is_End = true
	} else {
		ant.Path[ant.Pointeur].Ifant = true // sinon on dit juste que cette room est maintenant pleine
	}
	*tab_path_used = append(*tab_path_used, curentpath)          // on ajoute la path actuelle dans le tab_path_used
	return fmt.Sprint(ant.Name+ant.Path[ant.Pointeur].Name, " ") // fmt.Sprint() permet de convertir chaque élément en chaîne de caractère et de les concaténer de manière lisible (surtout si il y a plusieurs types)
	// équivalent à ça : return ant.Name + ant.Path[ant.Pointeur].Name
}

// Permet de savoir si une path est déjà dans notre tableau de paths
func Path_Not_Used(tab_path_used *[][2]*Room, curentpath [2]*Room) bool {
	// avec l'étoile, on modifie la valeur de l'élément pointé, alors que le & permet de récupérer l'adresse
	/*
		tab_path_used : (cette variable est un pointeur vers un tableau de tableaux)
			Lorsqu'il y a *[][2]*Room, cela signifie "un pointeur vers un tableau de tableaux de 2 pointeurs vers des objets Room"
		curentpath : (cette variable n'est pas un pointeur, c'est juste un tableau)
			Lorsqu'il y a [2]*Room, cela signifie "un tableau de 2 pointeurs vers des objets Room"

		Donc les 2 contiennent des pointeurs vers des objets Room mais les 2 ne sont pas des pointeurs
	*/
	for _, used_path := range *tab_path_used {
		if used_path == curentpath { // est ce que le curentpath est déjà dans le tab_path_used
			return false
		}
	}
	return true
}

/*
Chaque tableau de 2 Room dans tab_path_used représente un chemin que l'ant a déjà utilisé.
Le premier élément du tableau est la Room de départ du chemin et le deuxième élément est la Room d'arrivée du chemin.

Lorsque tab_path_used est passé à la fonction Next_Room, la fonction vérifie si le curentpath est déjà dans tab_path_used en utilisant la fonction path not used.
*/
