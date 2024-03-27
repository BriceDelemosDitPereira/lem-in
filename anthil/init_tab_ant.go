package anthil

import (
	"strconv"
)

// Permet d'attribuer une path à chaque fourmi de manière optimale
func Init_Tab_Ant(tab_path [][]*Room, nb_ant int) []Ant {
	var tab_ant []Ant
	for i := 1; i <= nb_ant; i++ {
		var a Ant
		a.Name = "L" + strconv.Itoa(i) + "-"
		tab_ant = append(tab_ant, a)
	}
	lengthmin := len(tab_path[0])
	pointer := 0
	for !Check_Ant_Path(tab_ant) {
		for _, path := range tab_path {
			if len(path) <= lengthmin && pointer < len(tab_ant) {
				tab_ant[pointer].Path = path
				pointer++
			}
		}
		lengthmin++
	}
	return tab_ant
}

// Permet de vérifier si path est remplie ou non dans la struct Ant
func Check_Ant_Path(tab_ant []Ant) bool {
	for _, ant := range tab_ant {
		if ant.Path == nil {
			return false
		}
	}
	return true
}
