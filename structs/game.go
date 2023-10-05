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
		c.KeyMap[0x0] = true
	} else {
		c.KeyMap[0x0] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		c.KeyMap[0x1] = true
	} else {
		c.KeyMap[0x1] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		c.KeyMap[0x2] = true
	} else {
		c.KeyMap[0x2] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		c.KeyMap[0x3] = true
	} else {
		c.KeyMap[0x3] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
		c.KeyMap[0x4] = true
	} else {
		c.KeyMap[0x4] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key5) {
		c.KeyMap[0x5] = true
	} else {
		c.KeyMap[0x5] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key6) {
		c.KeyMap[0x6] = true
	} else {
		c.KeyMap[0x6] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key7) {
		c.KeyMap[0x7] = true
	} else {
		c.KeyMap[0x7] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key8) {
		c.KeyMap[0x8] = true
	} else {
		c.KeyMap[0x8] = false
	}
	if ebiten.IsKeyPressed(ebiten.Key9) {
		c.KeyMap[0x9] = true
	} else {
		c.KeyMap[0x9] = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.KeyMap[0xA] = true
	} else {
		c.KeyMap[0xA] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		c.KeyMap[0xB] = true
	} else {
		c.KeyMap[0xB] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		c.KeyMap[0xC] = true
	} else {
		c.KeyMap[0xC] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.KeyMap[0xD] = true
	} else {
		c.KeyMap[0xD] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		c.KeyMap[0xE] = true
	} else {
		c.KeyMap[0xE] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		c.KeyMap[0xF] = true
	} else {
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
