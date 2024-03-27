package anthil

import "fmt"

func (fourmi *Fourmiliere) Research_Path() {
	var tabpath [][]*Room
	startroom, _ := fourmi.Get_Start() // On ne récupère plus le booléen mais la valeur
	// Il est possible de créer la fonction du dessous en dehors de cette fonction
	// Une fonction créé dans une fonction ne peut servir que dans cette fonction
	// La fonction juste en dessous va servir dans la prochaine fonction que l'on utilise dans cette fonction
	check_path := func(path []*Room, roomtesting *Room) bool { // On rentre dans la fonction un tableau de toutes les structures room et la valeur de chaque structure room
		for _, room := range path { // On regarde chaque room dans le tableau de toute les structures room
			if room == roomtesting { // On vérifie que la room que l'on check n'est pas déjà dans le chemin
				return false
			}
		}
		return true
	}
	// Pour le moment notre path []*Room est vide et cela va servir de buffer pour incrémenter notre tab_possibilty path
	var backtracking func(*Room, []*Room) // Go ne voulait pas initier la variable sur la ligne en dessous
	backtracking = func(curentroom *Room, path []*Room) {
		path = append(path, curentroom) // On ajoute la room actuelle après les room précédentes
		if curentroom.Type_Room == End {
			tabpath = append(tabpath, path) // Si on trouve la room End, alors on peut ajouter toute la path dans le tableau de path
			return                          // on peut arrêter la donc on sort de cette fonction anonyme avec return
		}
		for _, room := range curentroom.Link { // Après avoir trouvé les links, on va les range
			if check_path(path, room) {
				// Attention en manipulant les slices car go n'aime pas cela
				// On ne peut pas faire => temppath := []*Room(path[:]...)
				var temppath []*Room                 // buffer pour avoir tout le path avant de l'ajouter au tableau de path (tabpath)
				temppath = append(temppath, path...) // les ... permettent de mettre tout le path sans avoir besoin de préciser un index et le faire boucler
				// le buffer est nécessaire pour garder en mémoire les anciens path et revenir en arriève avec le contrôle de la fonction récursive (voir com plus bas)
				backtracking(room, temppath) // récursive où l'on renvoie la room lié à la précédente et notre path
				// cette récursive permet à la fois à continuer notre recherche de path mais aussi à faire marche arrière (backtracking) lorsqu'on ne peut plus avancer avec cette path
				// une des propriétés de la récursivité est que le contrôlé revient à l'execution de la dernière récursive si il n'y a pas d'issue à la dernière boucle for (voir com plus bas)
			}
		}
	}
	backtracking(startroom, []*Room{})
	fourmi.Tab_Possibility_Path = tabpath
}

// Permet de print les chemins
func (fourmi *Fourmiliere) Display_Path() {
	for index, path := range fourmi.Tab_Possibility_Path {
		// fmt.Println("path ", index+1, " : ", path)
		fmt.Print("path ", index+1, " : [ ")
		for _, room := range path {
			fmt.Print(room.Name, " ")
		}
		fmt.Println("]")
	}
}

/* Quand la fonction backtracking commence, il n'y a rien de rentré dedans donc elle ne va pas sur sa première ligne :
path = append(path, curentroom)

A la place, elle va directement à l'avant dernière de la fonction, c'est à dire:
backtracking(startroom, []*Room{})

ATTENTION pour voir la suite avec le debug, il faut impérativement mettre une point rouge au début de la fonction backtracking (backtracking = func(curentroom *Room, path []*Room) {)
C'est considéré comme une fonction à part entière donc pour rentrer dedans il faut ce point rouge.

Donc on ajoute la startroom dans le path. Par la suite ça ne rentre pas dans le if car le type de la curentroom n'est pas End.
Ensuite on rentre dans le range et il va parcourir le []Link dans la struct Room de notre curentroom. Après cela, il vérifie avec check_path que la pièce dans notre []Link n'est pas déjà dans notre path. Si c'est bon alors on retourne dans notre fonction backtracking et va mettre path dans un buffer. Ce buffer est nécessaire car il va permettre de garder en mémoire notre slice path qui sera utilisé par notre fonction récursive pour revenir en arrière si elle s'est trompée de chemin (voir la suite des commentaires). Donc on passe temppath dans la fonction backtracking pour garantir que chaque appel récursif a son propre chemin indépendant. Cela signifie aussi qu'à chaque récursive, la fonction a son propre path indépendant, ce qui permet de "revenir en arrière" sans perdre le chemin précédent. Cela peut être coûteux pour la mémoire (surtout avec beaucoup de Room et de chemin). Mais ces copies ne sont pas stockés indéfiniment. Chaque copie temporaire est éligible à la collecte des déchets (garbage collection) par le runtime de Go. Donc de la mémoire est libéré au fur et à mesure.

Que ce passe-t-il si le code ce trompe de chemin ? Alors la propriété de contrôle de la récursivité rentre en jeu. Lorsque la boucle for va se terminer sans trouver de chemin viable alors on va directement revenir à l'état de l'ancienne récursive et éventuellement trouver un nouveau chemin ou alors denouveau revenir à l'état encore antérieur de cette récursive. C'est la boucle for qui permet à la récursive de faire cet effet backtracking.
Donc la récursive avec la startroom ne sera jamais réutilisé après son utilisation au début car elle n'a pas de boucle for.

Avec un système de backtracking, on peut imaginer une boucle infinie mais ici le code est protégé par la fonction check_path qui vérifie si une room est dans le path pour éviter une boucle infinie. Cependant avec un input très grand, on peut imaginer que le code prendrait du temps pour arriver au bout et utiliser beaucoup de mémoire.

A chaque fois qu'une Room de type End est trouvée, alors on l'ajoute dans tabpath. Avec le return, on retourne dans le backtracking, et on enlève la room End en reprenant l'ancien path comme expliqué plus haut.
Et à la fin, quand tabpath aura tous les chemins. On l'ajoute dans notre structure Fourmiliere au sein de Tab_Possibility_path.*/
