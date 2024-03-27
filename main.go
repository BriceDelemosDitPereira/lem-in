package main

import (
	anthil "lem_in/anthil"
)

func main() {
	fourmi, check := anthil.Init_Data()
	if !check {
		return
	}
	fourmi.Research_Path()
	fourmi.Sort_Path()
	fourmi.Lot_Path()
	// fourmi.Display_Lot_Path()
	fourmi.Resolve()
}
