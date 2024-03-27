package anthil

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem_in/outils"
)

// On va récupérer l'input avec cette fonction et on va incrémenter notre structure Room (Add_Room, Add_Link, Check_Start_End)
func Init_Data() (Fourmiliere, bool) {
	arg := os.Args[1:]
	if len(arg) != 1 {
		outils.Print_err(outils.Err{Message: "Erreur : Veuillez entrer un seul d'argument valide"})
		return Fourmiliere{}, false
	}
	tabfile := outils.Readfile(arg[0]) // On récupère les éléments de notre premier argument et on en fait un tableau de byte
	var fourmi Fourmiliere             // une variable pour incrémenter notre structure
	var err error
	tabline := strings.Split(string(tabfile), "\n")     // Il est possible de convertir un tableau de byte en string
	fourmi.Start_Nb_Ant, err = strconv.Atoi(tabline[0]) // le nombre de fourmi au début est tout en haut
	outils.Check_err(err)
	if fourmi.Start_Nb_Ant <= 0 { // on vérifie si il y a un nombre de fourmis supérieur à 0
		outils.Print_err(outils.Err{Message: "Erreur : Vérifier le nombre de fourmis dans vos données"})
		return fourmi, false
	}
	var link_start_end [][]string
	// var count_link int
	for index := 1; index < len(tabline); index++ { // On va parcourir tout le tableau sauf à l'index 0 car c'est le nombre de fourmis (déjà pris avant)
		// En premier, on va récupérer la salle start et la salle end
		switch tabline[index] {
		case "##start":
			index++                                                       // Si on voit ##start alors on va à l'index suivant pour récupérer ses coordonnées
			fourmi.Add_Room(strings.Split(tabline[index], " ")[0], Start) // Pour appeler une méthode (une fonction avec comme récepteur une struct), il faut appeler la struct
			// C'est une fonction lié à notre struct
			// pour ne récupérer que le nom, on va split et ne récupérer que le premier élément
			continue // continue permet de passer à l'index suivant dans notre boucle
		case "##end":
			index++
			fourmi.Add_Room(strings.Split(tabline[index], " ")[0], End)
			continue
		}
		ispath := strings.Split(tabline[index], "-")
		if len(ispath) == 2 {
			link_start_end = append(link_start_end, ispath)
		}
		isroominfo := strings.Split(tabline[index], " ")
		switch {
		case len(ispath) == 2:
			if !fourmi.Add_Link(ispath) {
				outils.Print_err(outils.Err{
					Message: "Erreur : Vérifiez que vos liens entre vos salles sont bien sous ce format -> a-b",
					Reason:  tabline[index],
				})
			}
		case len(isroominfo) == 3:
			if !fourmi.Add_Room(isroominfo[0], Standard) {
				outils.Print_err(outils.Err{
					Message: "Erreur : Vérifiez que vous n'avez pas 2 salles avec le même nom",
					Reason:  tabline[index],
				})
				return fourmi, false
			}
		case tabline[index][0] == '#':
			continue
		default:
			outils.Print_err(outils.Err{
				Message: "Erreur : En dehors de ##start et ##end, il n'y a que 2 formats d'input pour ce code -> a-b || a b c",
				Reason:  tabline[index],
			})
			return fourmi, false
		}
	}
	if !fourmi.Check_Start_End() {
		outils.Print_err(outils.Err{
			Message: "Erreur : Attention il vous faut une salle start (##start) et une salle end (##end)",
		})
		return fourmi, false
	}

	if !fourmi.Check_Link_Start_End(link_start_end) {
		outils.Print_err(outils.Err{
			Message: "Erreur : Attention il vous faut un lien vers une salle start (##start) et une salle end (##end)",
		})
		return fourmi, false
	}
	fmt.Println(string(tabfile) + "\n")
	return fourmi, true
}
