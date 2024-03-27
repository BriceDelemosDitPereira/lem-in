package main

import (
	anthil "lem_in/anthil"
)

func main() {
	// Permet de récupérer les données dans la struct Fourmilière
	fourmi, check := anthil.Init_Data()
	if !check {
		return
	}
	// fmt.Println(fourmi)
	fourmi.Research_Path()
	// fourmi.Display_path()
	// fmt.Println("========================================")
	fourmi.Sort_Path()
	// fourmi.Display_Path()
	fourmi.Lot_Path()
	// fourmi.Display_Lot_Path()
	fourmi.Resolve()
}

/*func main() {
	var str string
	slice := [][]int{{9, 5}, {23, 3}, {16, 7}, {16, 3}, {16, 5}, {9, 3}, {1, 5}, {4, 8}}

	str += Map(slice)
	fmt.Print(str)
}

func Map(slice [][]int) string {
	var maxX, maxY int
	for i := range slice {
		if slice[i][0] > maxX {
			maxX = slice[i][0]
		}
		if slice[i][1] > maxY {
			maxY = slice[i][1]
		}
	}

	var str string

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			var isHere bool = false
			var nb int
			for i, k := range slice {
				if k[0] == x && k[1] == y {
					isHere = true
					nb = i
				}
			}

			if isHere {
				str += strconv.Itoa(nb)
			} else {
				str += " "
			}
		}
		str += "\n"
	}
	return str
} */
