package outils

import (
	"log"
	"os"
)

type Err struct {
	Message    string
	Reason     string
	Supplement string
}

// Permet de print l'erreur rencontrée
func Print_err(mess Err) {
	// if err != nil {
	// 	fmt.Println(err)
	// }
	if mess.Message != "" {
		log.Println(mess.Message)
	}
	if mess.Reason != "" {
		log.Println(mess.Reason)
	}
	if mess.Supplement != "" {
		log.Println(mess.Supplement)
	}
	os.Exit(1)
}

// Permet de Check si une erreur
func Check_err(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
