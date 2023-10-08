package structs

import "github.com/hajimehoshi/ebiten/audio"

type CPU struct {
	Registers   [16]byte
	PC          uint16
	I           uint16
	SP          uint16
	DT          uint16
	ST          uint16
	AudioPlayer *audio.Player
	Memory      [4096]byte
	Stack       [16]uint16
	Opcodes     []byte
	Screen      [64][32]byte
	KeyState    [16]byte

	//KeyMap map[int16]bool
}

func (c *CPU) InitCPU(data []byte) {
	c.PC = 0x200
	c.InitMemory(data)
	c.AddOpcodesToCPU()
	c.AudioPlayer = NewAudioPlayer()
}

func (c *CPU) InitMemory(data []byte) {
	c.AddROMToMemory(data)
	c.AddFontSetToMemory()
}

func (c *CPU) ClearScreen() {
	c.Screen = [64][32]byte{}
}

func (c *CPU) AddROMToMemory(data []byte) {
	for i, b := range data {
		c.Memory[0x200+i] = b
	}
}

func (c *CPU) AddFontSetToMemory() {
	fontset := []byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, //A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, //B
		0xF0, 0x80, 0x80, 0x80, 0xF0, //C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x0} // F
	for i, b := range fontset {
		c.Memory[i] = b
	}
}

//func StringToHexa(s string) int16 {
//	switch s {
//	case "Digit0":
//		return 0x0
//	case "Digit1":
//		return 0x1
//	case "Digit2":
//		return 0x2
//	case "Digit3":
//		return 0x3
//	case "Digit4":
//		return 0x4
//	case "Digit5":
//		return 0x5
//	case "Digit6":
//		return 0x6
//	case "Digit7":
//		return 0x7
//	case "Digit8":
//		return 0x8
//	case "Digit9":
//		return 0x9
//	case "A":
//		return 0xA
//	case "B":
//		return 0xB
//	case "C":
//		return 0xC
//	case "D":
//		return 0xD
//	case "E":
//		return 0xE
//	case "F":
//		return 0xF
//	default:
//		return 0x10
//	}
//}

//func (g *CPU) InitMapHexa() {
//
//	g.KeyMap = map[int16]bool{
//		0x0: false,
//		0x1: false,
//		0x2: falPlayer
//		0x3: false,
//		0x4: false,
//		0x5: false,
//		0x6: false,
//		0x7: false,
//		0x8: false,
//		0x9: false,
//		0xA: false,
//		0xB: false,
//		0xC: false,
//		0xD: false,
//		0xE: false,
//		0xF: false,
//	}
//}
