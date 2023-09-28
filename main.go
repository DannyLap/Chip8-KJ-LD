package main

import (
	"Chip8-JD/structs"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	cpu := new(structs.CPU)
	var memory [4096]byte
	cpu.Memory = memory

	file := os.Args[1]
	// Lit le contenu du fichier en mémoire
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}
	cpu.InitMemory(data)
	cpu.AddOpcodesToCPU()
	fmt.Println(cpu.Opcodes[0] * cpu.Opcodes[1])
	structs.OpenWindowEbiten() // remettre dans le main
}
