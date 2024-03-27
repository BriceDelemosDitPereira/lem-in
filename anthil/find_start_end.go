package anthil

// Permet de vérifier si il y a bien une Room Start et on envoi sa struct
func (fourmi *Fourmiliere) Get_Start() (*Room, bool) {
	for index, room := range fourmi.Anthil {
		if room.Type_Room == Start {
			return &fourmi.Anthil[index], true
		}
	}
	return nil, false
}

// Permet de vérifier si il y a bien une Room End et on envoi sa struct
func (fourmi *Fourmiliere) Get_End() (*Room, bool) {
	for index, room := range fourmi.Anthil {
		if room.Type_Room == End {
			return &fourmi.Anthil[index], true
		}
	}
	return nil, false
}

// Permet de regrouper le résultat du booleen Get_Start et Get_End
func (fourmi *Fourmiliere) Check_Start_End() bool {
	_, start := fourmi.Get_Start()
	_, end := fourmi.Get_End()
	return start && end
}
