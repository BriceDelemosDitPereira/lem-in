package anthil

type TypeRoom string

// Chaque donnée importante pour une fourmi
type Ant struct {
	Name     string
	Pointeur int
	Path     []*Room
	Is_End   bool
}

// Chaque donnée importante pour une pièce
type Room struct {
	Name      string
	Type_Room TypeRoom
	Ifant     bool
	Link      []*Room
}

// Chaque donnée importante pour la fourmilière
type Fourmiliere struct {
	Anthil               []Room
	Tab_Possibility_Path [][]*Room
	Lot_UniqueRoom_Path  [][][]*Room
	Start_Nb_Ant         int
}

// Permet d'incrémenter le nom et le type pour une structure Room
func (fourmi *Fourmiliere) Add_Room(nameRoom string, type_room TypeRoom) bool {
	var room Room
	for _, room := range fourmi.Anthil {
		if room.Name == nameRoom {
			return false
		}
	}
	room.Name = nameRoom
	room.Type_Room = type_room
	fourmi.Anthil = append(fourmi.Anthil, room)
	return true
}

// Permet d'incrémenter les liens entre 2 Room dans une structure Room (Link []*Room)
func (fourmi *Fourmiliere) Add_Link(tabroom []string) bool {
	if len(tabroom) != 2 || tabroom[0] == tabroom[1] {
		return false
	}
	var room1 *Room
	var room2 *Room
	for index, room := range fourmi.Anthil {
		if tabroom[0] == room.Name {
			room1 = &fourmi.Anthil[index]
		}
		if tabroom[1] == room.Name {
			room2 = &fourmi.Anthil[index]
		}
	}
	if room1 == room2 || room1 == nil || room2 == nil {
		return false
	}
	room1.Link = append(room1.Link, room2)
	room2.Link = append(room2.Link, room1)
	return true
}
