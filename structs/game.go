package structs

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Screen [64][32]byte
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{255, 0, 255, 255}
	whiteCol := color.RGBA{255, 255, 255, 255}
	//blackCol := color.RGBA{0, 0, 0, 255}
	//redCol := color.RGBA{255, 0, 0, 255}
	//blueCol := color.RGBA{0, 0, 255, 255}
	//greenCol := color.RGBA{0, 255, 0, 255}

	for i := 0; i < len(g.Screen); i++ {
		for j := 0; j < len(g.Screen[i]); j++ {
			if i == j {
				g.Screen[i][j] = 1
			}
		}
	}
	for i := len(g.Screen) - 1; i > 0; i-- {
		for j := len(g.Screen[i]) - 1; j > 0; j-- {
			if i == j {
				g.Screen[i][j] = 1
			}
		}
	}

	//g.Screen[63][31] = 1
	//g.Screen[63][0] = 1
	//g.Screen[0][31] = 1
	//g.Screen[0][0] = 1
	//DrawASquare(63*5, 31*5, screen, 5, redCol)
	//DrawASquare(63*5, 0*5, screen, 5, blueCol)
	//DrawASquare(0*5, 31*5, screen, 5, greenCol)
	//DrawASquare(0*5, 0*5, screen, 5, purpleCol)

	for i := 0; i < len(g.Screen); i++ {
		for j := 0; j < len(g.Screen[i]); j++ {
			if g.Screen[i][j] == 1 {
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func OpenWindowEbiten() {
	ebiten.SetWindowSize(1000, 900)
	ebiten.SetWindowTitle("Chip 8")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
