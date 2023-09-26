package main

import (
	"Chip8-JD/structs"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	cpu := new(structs.CPU)
	file := os.Args[1]
	// Lit le contenu du fichier en mémoire
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	cpu.Memory = data

	structs.OpenWindowEbiten()
}
