package anthil

import "fmt"

// Permet de trouver chaque chemin commencant par une Room Start et une Room End
func (fourmi *Fourmiliere) Research_Path() {
	var tabpath [][]*Room
	startroom, _ := fourmi.Get_Start()
	check_path := func(path []*Room, roomtesting *Room) bool {
		for _, room := range path {
			if room == roomtesting {
				return false
			}
		}
		return true
	}

	var backtracking func(*Room, []*Room)
	backtracking = func(curentroom *Room, path []*Room) {
		path = append(path, curentroom)
		if curentroom.Type_Room == End {
			tabpath = append(tabpath, path)
			return
		}
		for _, room := range curentroom.Link {
			if check_path(path, room) {
				var temppath []*Room
				temppath = append(temppath, path...)
				backtracking(room, temppath)
			}
		}
	}
	backtracking(startroom, []*Room{})
	fourmi.Tab_Possibility_Path = tabpath
}

// Permet de print les chemins
func (fourmi *Fourmiliere) Display_Path() {
	for index, path := range fourmi.Tab_Possibility_Path {
		fmt.Print("path ", index+1, " : [ ")
		for _, room := range path {
			fmt.Print(room.Name, " ")
		}
		fmt.Println("]")
	}
}
