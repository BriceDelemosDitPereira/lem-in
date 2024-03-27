package anthil

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem_in/outils"
)

// Permet de récupérer les données de notre input
func Init_Data() (Fourmiliere, bool) {
	arg := os.Args[1:]
	if len(arg) != 1 {
		outils.Print_err(outils.Err{Message: "ERROR: invalid data format, invalid number of arguments (need only 1 argument)"})
		return Fourmiliere{}, false
	}
	tabfile := outils.Readfile(arg[0])
	var fourmi Fourmiliere
	var err error
	tabline := strings.Split(string(tabfile), "\n")
	fourmi.Start_Nb_Ant, err = strconv.Atoi(tabline[0])
	outils.Check_err(err)
	if fourmi.Start_Nb_Ant <= 0 {
		outils.Print_err(outils.Err{Message: "ERROR: invalid data format, invalid number of Ants"})
		return fourmi, false
	}
	var link_start_end [][]string
	for index := 1; index < len(tabline); index++ {
		switch tabline[index] {
		case "##start":
			index++
			fourmi.Add_Room(strings.Split(tabline[index], " ")[0], Start)
			continue
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
					Message: "ERROR: invalid data format, used this format for links rooms -> a-b",
					Reason:  tabline[index],
				})
				return fourmi, false
			}
		case len(isroominfo) == 3:
			if !fourmi.Add_Room(isroominfo[0], Standard) {
				outils.Print_err(outils.Err{
					Message: "ERROR: invalid data format, each room needs its own name",
					Reason:  tabline[index],
				})
				return fourmi, false
			}
		case tabline[index][0] == '#':
			continue
		default:
			outils.Print_err(outils.Err{
				Message: "ERROR: invalid data format, except ##start and ##end, only 2 formats -> a-b || a b c",
				Reason:  tabline[index],
			})
			return fourmi, false
		}
	}
	if !fourmi.Check_Start_End() {
		outils.Print_err(outils.Err{
			Message: "ERROR: invalid data format, need one room start (##start) and one room end (##end)",
		})
		return fourmi, false
	}
	if !fourmi.Check_Link_Start_End(link_start_end) {
		outils.Print_err(outils.Err{
			Message: "ERROR: invalid data format, no link for the room start and/or end",
		})
		return fourmi, false
	}
	fmt.Println(string(tabfile) + "\n")
	return fourmi, true
}
