package structs

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// type Game struct {
// 	Registers []byte
// 	PC        uint16
// 	I         uint16
// 	SP        uint16
// 	DT        uint16
// 	ST        uint16
// 	Memory    [4096]byte
// 	Stack     [16]uint16
// 	Opcodes   []byte
// 	Screen    [64][32]int
// }

//var (
//	keyStates = make(map[ebiten.Key]bool)
//)

func (c *CPU) Update() error {
	c.OpcodesReading()

	if ebiten.IsKeyPressed(ebiten.Key0) {
		c.KeyState[0] = 1
		c.KeyMap[0x0] = true
	} else {
		c.KeyState[0] = 0
		c.KeyMap[0x0] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		c.KeyState[1] = 1
		c.KeyMap[0x1] = true
	} else {
		c.KeyState[1] = 0
		c.KeyMap[0x1] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		c.KeyState[2] = 1
		c.KeyMap[0x2] = true
	} else {
		c.KeyState[2] = 0
		c.KeyMap[0x2] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		c.KeyState[3] = 1
		c.KeyMap[0x3] = true
	} else {
		c.KeyState[3] = 0
		c.KeyMap[0x3] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
		c.KeyState[4] = 1
		c.KeyMap[0x4] = true
	} else {
		c.KeyState[4] = 0
		c.KeyMap[0x4] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key5) {
		c.KeyState[5] = 1
		c.KeyMap[0x5] = true
	} else {
		c.KeyState[5] = 0
		c.KeyMap[0x5] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key6) {
		c.KeyState[6] = 1
		c.KeyMap[0x6] = true
	} else {
		c.KeyState[6] = 0
		c.KeyMap[0x6] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key7) {
		c.KeyState[7] = 1
		c.KeyMap[0x7] = true
	} else {
		c.KeyState[7] = 0
		c.KeyMap[0x7] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key8) {
		c.KeyState[8] = 1
		c.KeyMap[0x8] = true
	} else {
		c.KeyState[8] = 0
		c.KeyMap[0x8] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key9) {
		c.KeyState[9] = 1
		c.KeyMap[0x9] = true
	} else {
		c.KeyState[9] = 0
		c.KeyMap[0x9] = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.KeyState[0xA] = 1
		c.KeyMap[0xA] = true
	} else {
		c.KeyState[0xA] = 0
		c.KeyMap[0xA] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		c.KeyState[0xB] = 1
		c.KeyMap[0xB] = true
	} else {
		c.KeyState[0xB] = 0
		c.KeyMap[0xB] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		c.KeyState[0xC] = 1
		c.KeyMap[0xC] = true
	} else {
		c.KeyState[0xC] = 0
		c.KeyMap[0xC] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.KeyState[0xD] = 1
		c.KeyMap[0xD] = true
	} else {
		c.KeyState[0xD] = 0
		c.KeyMap[0xD] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		c.KeyState[0xE] = 1
		c.KeyMap[0xE] = true
	} else {
		c.KeyState[0xE] = 0
		c.KeyMap[0xE] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		c.KeyState[0xF] = 1
		c.KeyMap[0xF] = true
	} else {
		c.KeyState[0xF] = 0
		c.KeyMap[0xF] = false
	}

	for key, value := range c.KeyMap {
		if value {
			fmt.Println("La touche qui a pour valeur : ", key, " est enfoncée")
		}
	}

	return nil
}

func (c *CPU) Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{255, 0, 255, 255}
	whiteCol := color.RGBA{255, 255, 255, 255}
	//blackCol := color.RGBA{0, 0, 0, 255}
	//redCol := color.RGBA{255, 0, 0, 255}
	//blueCol := color.RGBA{0, 0, 255, 255}
	//greenCol := color.RGBA{0, 255, 0, 255}

	for i := 0; i < len(c.Screen); i++ {
		for j := 0; j < len(c.Screen[i]); j++ {
			if c.Screen[i][j] == 1 {
				DrawASquare(i*5, j*5, screen, 5, whiteCol)
			} else {
				DrawASquare(i*5, j*5, screen, 5, purpleCol)
			}
		}
	}
}

func DrawASquare(x int, y int, screen *ebiten.Image, size int, color color.RGBA) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			screen.Set(x+i, y+j, color)
		}
	}
}

func (c *CPU) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func OpenWindowEbiten(cpu *CPU) {
	ebiten.SetWindowSize(1000, 900)
	ebiten.SetWindowTitle("Chip 8")
	if err := ebiten.RunGame(cpu); err != nil {
		log.Fatal(err)
	}
}
