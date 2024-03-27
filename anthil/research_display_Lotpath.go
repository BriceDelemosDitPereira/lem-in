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
			// On va comparer la path actuel avec d'autres path dans Tab_Possibility_path
			if Compare(use_room, fourmi.Tab_Possibility_Path[index_bis]) { // Si il n'y a aucune room en commun dans les 2 path
				temp_lot_path = append(temp_lot_path, fourmi.Tab_Possibility_Path[index_bis]) // on ajoute cette path dans notre buffer
				use_room = append(use_room, fourmi.Tab_Possibility_Path[index_bis]...)        // On ajoute aussi dans use_room car il ne faudra pas avoir de salle en commun avec les 2 path à l'intérieur
			}
		}
		fourmi.Lot_UniqueRoom_Path = append(fourmi.Lot_UniqueRoom_Path, temp_lot_path) // Pour finir on incrémente la struct Fourmiliere
	}
}

// Permet de vérifier si il y a des room similaires au sein de 2 path différentes
func Compare(listroom []*Room, path []*Room) bool {
	// On range la path du Tab_Possibility_path
	for _, room := range path {
		if room.Type_Room == Start || room.Type_Room == End { // On passe cette étape car ces rooms sont similaires dans toutes les path
			continue // Le continue permet de passer à l'index suivant dans la boucle for
		}
		// On range la path que l'on check
		for index := range listroom {
			if room == listroom[index] { // On compare les 2
				return false
			}
		}
	}
	return true
	// L'objectif est de trouver plusieurs path avec aucune room en commun
}

// Permet de print les chemins qui n'utilisent aucune room en commun (regroupé sous forme de lot)
func (fourmi *Fourmiliere) Display_Lot_Path() {
	fmt.Println("======================================================")
	for index, lotpath := range fourmi.Lot_UniqueRoom_Path { // Pour print par lot
		fmt.Print("lot ", index+1, ":")
		for _, path := range lotpath { // Pour print par chemin
			// fmt.Println("	", path)
			fmt.Print("		[")
			for _, room := range path { // Pour print par room
				fmt.Print(room.Name, " ")
			}
			fmt.Println("]")
		}
		fmt.Println("======================================================")
	}
}
