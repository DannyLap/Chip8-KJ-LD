package structs

import (
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

var (
	keyStates = make(map[ebiten.Key]bool)
)

func (c *CPU) Update() error {
	c.OpcodesReading()
	return nil

	//if ebiten.IsKeyPressed(ebiten.KeyA) {
	//	keyStates[ebiten.KeyA] = true
	//} else {
	//	keyStates[ebiten.KeyA] = false
	//}
	//
	//if keyStates[ebiten.KeyA] {
	//	fmt.Println("La touche 'A' est enfoncée.")
	//}
	//
	//return nil
}

func (c *CPU) Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{255, 0, 255, 255}
	whiteCol := color.RGBA{255, 255, 255, 255}
	//blackCol := color.RGBA{0, 0, 0, 255}
	//redCol := color.RGBA{255, 0, 0, 255}
	//blueCol := color.RGBA{0, 0, 255, 255}
	//greenCol := color.RGBA{0, 255, 0, 255}

	// for i := 0; i < len(g.Screen); i++ {
	// 	for j := 0; j < len(g.Screen[i]); j++ {
	// 		if i == j {
	// 			g.Screen[i][j] = 1
	// 		}
	// 	}
	// }
	// for i := len(g.Screen) - 1; i > 0; i-- {
	// 	for j := len(g.Screen[i]) - 1; j > 0; j-- {
	// 		if i == j {
	// 			g.Screen[i][j] = 1
	// 		}
	// 	}
	// }

	//g.Screen[63][31] = 1
	//g.Screen[63][0] = 1
	//g.Screen[0][31] = 1
	//g.Screen[0][0] = 1
	//DrawASquare(63*5, 31*5, screen, 5, redCol)
	//DrawASquare(63*5, 0*5, screen, 5, blueCol)
	//DrawASquare(0*5, 31*5, screen, 5, greenCol)
	//DrawASquare(0*5, 0*5, screen, 5, purpleCol)

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
