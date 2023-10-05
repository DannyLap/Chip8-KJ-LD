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

	//if ebiten.IsKeyPressed(ebiten.KeyA) {
	//	keyStates[ebiten.KeyA] = true
	//	fmt.Println("Yesssss")
	//} else {
	//	keyStates[ebiten.KeyA] = false
	//	fmt.Println("No")
	//}
	//
	//if keyStates[ebiten.KeyA] {
	//	fmt.Println("La touche 'A' est enfoncée.")
	//}

	if ebiten.IsKeyPressed(ebiten.Key0) || ebiten.IsKeyPressed(ebiten.KeyNumpad0) {
		c.Key = ebiten.KeyDigit0.String()
	} else if ebiten.IsKeyPressed(ebiten.Key1) || ebiten.IsKeyPressed(ebiten.KeyNumpad1) {
		c.Key = ebiten.KeyDigit1.String()
	} else if ebiten.IsKeyPressed(ebiten.Key2) || ebiten.IsKeyPressed(ebiten.KeyNumpad2) {
		c.Key = ebiten.KeyDigit2.String()
	} else if ebiten.IsKeyPressed(ebiten.Key3) || ebiten.IsKeyPressed(ebiten.KeyNumpad3) {
		c.Key = ebiten.KeyDigit3.String()
	} else if ebiten.IsKeyPressed(ebiten.Key4) || ebiten.IsKeyPressed(ebiten.KeyNumpad4) {
		c.Key = ebiten.KeyDigit4.String()
	} else if ebiten.IsKeyPressed(ebiten.Key5) || ebiten.IsKeyPressed(ebiten.KeyNumpad5) {
		c.Key = ebiten.KeyDigit5.String()
	} else if ebiten.IsKeyPressed(ebiten.Key6) || ebiten.IsKeyPressed(ebiten.KeyNumpad6) {
		c.Key = ebiten.KeyDigit6.String()
	} else if ebiten.IsKeyPressed(ebiten.Key7) || ebiten.IsKeyPressed(ebiten.KeyNumpad7) {
		c.Key = ebiten.KeyDigit7.String()
	} else if ebiten.IsKeyPressed(ebiten.Key8) || ebiten.IsKeyPressed(ebiten.KeyNumpad8) {
		c.Key = ebiten.KeyDigit8.String()
	} else if ebiten.IsKeyPressed(ebiten.Key9) || ebiten.IsKeyPressed(ebiten.KeyNumpad9) {
		c.Key = ebiten.KeyDigit9.String()
	} else if ebiten.IsKeyPressed(ebiten.KeyQ) { // petit soucis, ebiten comprends le clavier comme du querty donc le Q vaut A
		c.Key = ebiten.KeyA.String()
	} else if ebiten.IsKeyPressed(ebiten.KeyB) {
		c.Key = ebiten.KeyB.String()
	} else if ebiten.IsKeyPressed(ebiten.KeyC) {
		c.Key = ebiten.KeyC.String()
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.Key = ebiten.KeyD.String()
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		c.Key = ebiten.KeyE.String()
	} else if ebiten.IsKeyPressed(ebiten.KeyF) {
		c.Key = ebiten.KeyF.String()
	}

	if c.Key != "" {
		fmt.Println("Value of c.Key = ", c.Key)
	}

	//c.Key = ""

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
