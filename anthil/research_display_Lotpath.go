package anthil

import "fmt"

// Permet de regrouper les chemins si ils n'ont aucune room en commun
func (fourmi *Fourmiliere) Lot_Path() {
	for index, path := range fourmi.Tab_Possibility_Path {
		var use_room []*Room
		var temp_lot_path [][]*Room
		temp_lot_path = append(temp_lot_path, path)
		use_room = append(use_room, path...)
		for index_bis := index + 1; index_bis < len(fourmi.Tab_Possibility_Path); index_bis++ {
			if Compare(use_room, fourmi.Tab_Possibility_Path[index_bis]) {
				temp_lot_path = append(temp_lot_path, fourmi.Tab_Possibility_Path[index_bis])
				use_room = append(use_room, fourmi.Tab_Possibility_Path[index_bis]...)
			}
		}
		fourmi.Lot_UniqueRoom_Path = append(fourmi.Lot_UniqueRoom_Path, temp_lot_path)
	}
}

// Permet de vérifier si il y a des room similaires au sein de 2 path différentes
func Compare(listroom []*Room, path []*Room) bool {
	for _, room := range path {
		if room.Type_Room == Start || room.Type_Room == End {
			continue
		}
		for index := range listroom {
			if room == listroom[index] {
				return false
			}
		}
	}
	return true
}

// Permet de print les chemins qui n'utilisent aucune room en commun (regroupé sous forme de lot)
func (fourmi *Fourmiliere) Display_Lot_Path() {
	fmt.Println("======================================================")
	for index, lotpath := range fourmi.Lot_UniqueRoom_Path {
		fmt.Print("lot ", index+1, ":")
		for _, path := range lotpath {
			fmt.Print("		[")
			for _, room := range path {
				fmt.Print(room.Name, " ")
			}
			fmt.Println("]")
		}
		fmt.Println("======================================================")
	}
}
