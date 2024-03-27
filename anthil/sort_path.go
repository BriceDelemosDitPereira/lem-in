package anthil

// Permet de classer les chemins par rapport à leurs tailles (ordre croissant)
func (fourmi *Fourmiliere) Sort_Path() {
	for index := 1; index < len(fourmi.Tab_Possibility_Path); index++ {
		if len(fourmi.Tab_Possibility_Path[index-1]) > len(fourmi.Tab_Possibility_Path[index]) {
			temp := fourmi.Tab_Possibility_Path[index]
			fourmi.Tab_Possibility_Path[index] = fourmi.Tab_Possibility_Path[index-1]
			fourmi.Tab_Possibility_Path[index-1] = temp
			index = 0
		}
	}
}
