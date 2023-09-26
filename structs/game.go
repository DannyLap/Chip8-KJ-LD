package structs

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func OpenWindowEbiten() {
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Chip 8")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
