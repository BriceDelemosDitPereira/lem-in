package anthil

type TypeRoom string

/* Room est de type struct, on va pouvoir appeler et incrémenter chaque élément en son sein
Chaque élément dans une struct est lié et ensuite on peut faire un tableau de struct pour parcourir chaque struct
Par exemple artist1 sera lié à nom1 et ils seront dans la même structure
Dans la structure 2 on aura artist2 et nom2 */

type Room struct {
	Name      string
	Type_Room TypeRoom
	// ne sert pas avec le code actuel
	// Nbant int
	Ifant bool    // Permet de savoir si une fourmi est présente dans la Room
	Link  []*Room // Permet de lier certaine struct room à d'autres
}

type Fourmiliere struct {
	Anthil               []Room
	Tab_Possibility_Path [][]*Room
	Lot_UniqueRoom_Path  [][][]*Room
	Start_Nb_Ant         int
}

type Ant struct {
	Name     string
	Pointeur int // Permet de cibler une structure Ant (donc une fourmi parmi celle que l'on a)
	Path     []*Room
	Is_End   bool
}

// Permet d'incrémenter le nom et le type pour une structure Room
func (fourmi *Fourmiliere) Add_Room(nameRoom string, type_room TypeRoom) bool {
	var room Room
	// Permet de vérifier qu'il n'y a pas 2 fois le même nom pour une Room
	for _, room := range fourmi.Anthil {
		if room.Name == nameRoom {
			return false
		}
	}
	room.Name = nameRoom
	room.Type_Room = type_room
	fourmi.Anthil = append(fourmi.Anthil, room) // on incrémente anthil avec le name de room et le typ_room de room
	// On peut le faire en une ligne:
	// fourmi.Anthil = append(fourmi.Anthil, Room{Name: nameRoom, Type_room: type_room})
	return true
}

// Permet d'incrémenter les liens dans une structure Room (Link []*Room)
func (fourmi *Fourmiliere) Add_Link(tabroom []string) bool {
	/* un lien aura le format a-b et pour faire notre tableau, on va split avec "-" donc notre length ne sera que de 2 */
	if len(tabroom) != 2 || tabroom[0] == tabroom[1] { // un lien ne peut être lié à la même salle
		return false
	}
	var room1 *Room
	var room2 *Room
	for index, room := range fourmi.Anthil {
		if tabroom[0] == room.Name { // Permet de vérifier si une pièce porte ce nom
			room1 = &fourmi.Anthil[index] // On attribue sa valeur à room1
		}
		if tabroom[1] == room.Name {
			room2 = &fourmi.Anthil[index]
		}
	}
	if room1 == room2 || room1 == nil || room2 == nil { // room1 et 2 ne peuvent être la même pièce et il ne peut y avoir de lien avec rien
		return false
	}

	room1.Link = append(room1.Link, room2) // On va lier la structure room1 avec la structure room2
	room2.Link = append(room2.Link, room1) // On va aussi lier dans l'autre sens car les chemins peuvent aller dans les 2 sens
	return true
}
