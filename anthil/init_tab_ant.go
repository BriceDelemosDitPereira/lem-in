package anthil

import (
	"strconv"
)

func Init_Tab_Ant(tab_path [][]*Room, nb_ant int) []Ant { // On met dans la fonction le Tab_Possibility_Path de la struct Room et le Start_Nb_Ant
	var tab_ant []Ant
	for i := 1; i <= nb_ant; i++ { // On commence à 1 car il ne peut pas y avoir 0 fourmi
		var a Ant                            // obligé d'initier une variable de type Ant pour incrémenter Ant                          // On initie une varibale de type Ant, donc avec les éléments présents dans la struct Ant
		a.Name = "L" + strconv.Itoa(i) + "-" // On incrémente Name au sein de la struct Ant, on ajoute le numéro de la fourmi et on ajoute le tiret
		tab_ant = append(tab_ant, a)
	}
	lengthmin := len(tab_path[0])  // On récupère la len du plus petit path (car déjà rangé dans l'ordre croissant)
	pointer := 0                   // va permettre d'incrémenter les path de chaque fourmi, car on va l'augmenter dès qu'une fourmi à une path != nil
	for !Check_Ant_Path(tab_ant) { // on vérifie que Ant.Path est égal à nil car on veut l'incrémenter
		for _, path := range tab_path {
			if len(path) <= lengthmin && pointer < len(tab_ant) { // len de la path doit être inférieur à la lengthmin et pointer doit être inférieur au nombre de fourmi
				// len(path) <= lengthmin va permettre d'optimiser le nombre de fourmis par path (voir commentaires en dessous)
				// pointer < len(tab_ant) dès que le pointer sera égal au nombre de fourmi alors on arrète la boucle, si on fait commencer le pointer
				// obligé de faire commencer le pointer à 0 car sinon on a un overflow au niveau du tableau dans la struct Ant (overflow si on fait pointer = 1 et pointer <= len(tab_ant)
				/*fmt.Println("-------------------------------------")
				fmt.Println(tab_ant)
				fmt.Println("-------------------------------------")
				fmt.Println(len(path))
				fmt.Println(path)
				fmt.Println(lengthmin)
				fmt.Println(pointer)
				fmt.Println("-------------------------------------")*/
				tab_ant[pointer].Path = path // cela va permettre de donner une path à chaque fourmi, car on va attribuer une path pour chacune d'entre elle
				pointer++
			}
		}
		lengthmin++
	}
	// Permet d'afficher les éléments que l'on a rentré
	//for _, ant := range tab_ant {
	//	fmt.Println(ant.Name, ant.Path) // On peut faire apparaître autant d'élément que l'on souhaite de la struct Ant, là on a affiché Name et Path
	//}
	return tab_ant
}

// Permet de vérifier si path est remplie ou non dans la struct Ant
func Check_Ant_Path(tab_ant []Ant) bool {
	for _, ant := range tab_ant {
		if ant.Path == nil {
			return false
		}
	}
	return true
}

/* On va expliquer le len(path) <= lengthmin :
Imaginons 2 path pour 4 fourmis, une de 2 chambres et une de 4 chambres. Au début lengthmin sera égal à 2. On va parcourir nos path avec notre range de tab_path.
la première path a un len de 2 donc on ajoute cette path pour la fourmi 1 et ensuite le range va sur la 2ème path et elle est trop grande donc on sort de la boucle et on incrémente lenghtmin
lengthmin = 3, on peut denouveau rajouter une fourmis dans la path 1 mais toujours pas dans la path 2.
lengthmin = 4, on peut ajouter une fourmi dans la path 1 et dans la path 2 => cela optimise donc le temps de notre programme
(pour comprendre cela il faut tester avec uniquement Tab_Possibility_Path et pas avec le Lot)

Maintenant il est encore plus optimisé avec un système de lot. Il suffit pour cela de faire un range dans le Lot_UniqueRomm_Path et de faire rentrer chaque élément dans ce fichier golang.

*/
