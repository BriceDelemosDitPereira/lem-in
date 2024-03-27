package anthil

// Permet de classer les chemins par rapport à leuurs tailles (du plus petit au plus grand)
func (fourmi *Fourmiliere) Sort_Path() {
	for index := 1; index < len(fourmi.Tab_Possibility_Path); index++ { // on continue la boucle tant qu'il y a des chemins
		if len(fourmi.Tab_Possibility_Path[index-1]) > len(fourmi.Tab_Possibility_Path[index]) { // si la route à l'index-1 est plus grande que la route à l'index
			temp := fourmi.Tab_Possibility_Path[index] // Buffer
			fourmi.Tab_Possibility_Path[index] = fourmi.Tab_Possibility_Path[index-1]
			fourmi.Tab_Possibility_Path[index-1] = temp
			index = 0 // On repart à l'index 0 pour rechecker
		}
	}
}
