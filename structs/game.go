package structs

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenT [32][64]byte
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{255, 0, 255, 255}
	for i := 0; i < len(g.ScreenT); i++ {
		for j := 0; j < len(g.ScreenT[i]); j++ {
			if i == j {
				g.ScreenT[i][j] = 1
			}
		}
	}

	for i := 0; i < len(g.ScreenT); i++ {
		for j := 0; j < len(g.ScreenT[i]); j++ {
			if g.ScreenT[i][j] == 1 {
				screen.Set(i*10, j*10, purpleCol)
			}
		}
	}
	//ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func OpenWindowEbiten() {
	//var screen [32][64]byte
	//for i := 0; i < len(screen); i++ {
	//	for j := 0; j < len(screen[i]); j++ {
	//		if i == j {
	//			screen[i][j] = 1
	//		}
	//	}
	//	//fmt.Print(i)
	//	//fmt.Println(screen[i])
	//}

	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Chip 8")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
