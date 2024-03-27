package anthil

import (
	"fmt"
)

func (fourmi *Fourmiliere) Resolve() {
	var Printfinal string
	var nb_line_final int
	for _, lot := range fourmi.Lot_UniqueRoom_Path {
		// fmt.Println(lot)
		// fmt.Println("===========================")
		tabAnt := Init_Tab_Ant(lot, fourmi.Start_Nb_Ant) // on va donner un nom à chaque fourmi et attribuer une path
		tempstr := ""
		compt := 0
		for !Ant_At_End(tabAnt) { // tant que la fourmi n'est pas dans la Room End
			compt++
			tab_path_used := [][2]*Room{}
			for index := range tabAnt {
				tempstr += tabAnt[index].Next_Room(&tab_path_used)
				// Pour chaque fourmi, la méthode Next_Room est appelée pour déplacer la fourmi à la room suivante sur son chemin
			}
			tempstr += "\n" // Permet un saut à la ligne entre chaque étape
		}
		// fmt.Println(tempstr)
		// fmt.Println("-----------------------------------------")
		if nb_line_final == 0 || nb_line_final > compt { // Permet de garder une trace du nombre minimal de lignes nécessaires pour atteindre la fin
			/*
				Après que toutes les fourmis aient atteint la fin, cette condition vérifie si le nombre de lignes du chemin actuel (compt) est inférieur au nombre de lignes du chemin le plus court trouvé jusqu'à présent (nb_line_final).
				Si c'est le cas, alors nb_line_final et Printfinal sont mis à jour avec le nombre de lignes et la représentation textuelle du chemin actuel
				C'est une sorte de backtracking
			*/
			nb_line_final = compt
			Printfinal = tempstr
		}
	}
	fmt.Print(Printfinal)
	fmt.Printf("\nNumber of turns: %v\n", nb_line_final)
}

// Permet de savoir si une fourmi est dans la Room End
func Ant_At_End(tab_ant []Ant) bool {
	for _, ant := range tab_ant {
		if !ant.Is_End { // on retourne false si Ant n'est pas à la fin
			return false
		}
	}
	return true
}

/*
J'avais pour idée de créer un autre dossier pour la structure Ant mais ce n'est pas possible pour 2 raisons liées :
	- Le cycle d'import illégal :
		Le résolveur ne peut pas être mis dans le dossier anthil si il y a un autre dossier avec la struct Ant car le dossier avec la structure Ant
		et le dossier avec la structure Fourmilière aurait tout les 2 besoins de l'autre.
		Cela ferait anthil -> ant -> anthil et donc c'est un cycle d'import qui n'est pas permis dans go (pour éviter les boucles)
		Donc il est très import de bien réfléchir à la structure du code avant
	- Impossible de faire une nouvelle méthode dans un autre dossier :
		Il est nécessaire d'être dans le dossier de la structure en question pour lui créer une méthode donc le resolveur doit forcément être dans le dossier anthil
		car c'est une méthode de la structure Fourmiliere
=> 2 raisons qui font que je ne peux faire ce code qu'avec un dossier unique de fonction
*/
