package anthil

import (
	"fmt"
)

// Permet de résoudre lem-in
func (fourmi *Fourmiliere) Resolve() {
	var Printfinal string
	var nb_line_final int
	for _, lot := range fourmi.Lot_UniqueRoom_Path {
		tabAnt := Init_Tab_Ant(lot, fourmi.Start_Nb_Ant)
		tempstr := ""
		compt := 0
		for !Ant_At_End(tabAnt) {
			compt++
			tab_path_used := [][2]*Room{}
			for index := range tabAnt {
				tempstr += tabAnt[index].Next_Room(&tab_path_used)
			}
			tempstr += "\n"
		}
		if nb_line_final == 0 || nb_line_final > compt {
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
		if !ant.Is_End {
			return false
		}
	}
	return true
}
