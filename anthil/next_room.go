package anthil

import (
	"fmt"
)

// Permet de déplacer la fourmi dans la salle suivante
func (ant *Ant) Next_Room(tab_path_used *[][2]*Room) string {
	if ant.Is_End {
		return ""
	}
	curentpath := [2]*Room{ant.Path[ant.Pointeur], ant.Path[ant.Pointeur+1]}
	if !Path_Not_Used(tab_path_used, curentpath) || ant.Path[ant.Pointeur+1].Ifant {
		return ""
	}
	ant.Path[ant.Pointeur].Ifant = false
	ant.Pointeur++
	if ant.Path[ant.Pointeur].Type_Room == End {
		ant.Is_End = true
	} else {
		ant.Path[ant.Pointeur].Ifant = true
	}
	*tab_path_used = append(*tab_path_used, curentpath)
	return fmt.Sprint(ant.Name+ant.Path[ant.Pointeur].Name, " ")
}

// Permet de savoir si une path est déjà dans notre tableau de path
func Path_Not_Used(tab_path_used *[][2]*Room, curentpath [2]*Room) bool {
	for _, used_path := range *tab_path_used {
		if used_path == curentpath {
			return false
		}
	}
	return true
}
