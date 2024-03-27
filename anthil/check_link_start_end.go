package anthil

// Permet de savoir si il y a un lien entre la Room Start ou End et une autre Room
func (fourmi *Fourmiliere) Check_Link_Start_End(link_start_end [][]string) bool {
	var room1 *Room
	var room2 *Room

	var count_Start int
	var count_End int

	for _, path := range link_start_end {
		for index, room := range fourmi.Anthil {
			if path[0] == room.Name {
				room1 = &fourmi.Anthil[index]
			}
			if path[1] == room.Name {
				room2 = &fourmi.Anthil[index]
			}
		}
		if room1.Type_Room == Start && room2.Type_Room == Standard || room1.Type_Room == Standard && room2.Type_Room == Start {
			count_Start++
		}
		if room1.Type_Room == End && room2.Type_Room == Standard || room1.Type_Room == Standard && room2.Type_Room == End {
			count_End++
		}
		if room1.Type_Room == Start && room2.Type_Room == End || room1.Type_Room == End && room2.Type_Room == Start {
			count_Start++
			count_End++
		}
	}
	if count_Start == 0 || count_End == 0 {
		return false
	}
	return true
}
