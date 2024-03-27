package outils

import "os"

// Permet de lire l'input
func Readfile(path string) []byte {
	file, err := os.ReadFile(path)
	Check_err(err)
	return file
}
