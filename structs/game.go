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
	cyclesPerFrame := 20

	for i := 0; i < cyclesPerFrame; i++ {
		c.KeyPress()
		c.OpcodesReading()
	}

	if c.SoundEnabled() {
		c.AudioPlayer.SetVolume(1)
	} else {
		c.AudioPlayer.SetVolume(0)
	}

	c.UpdateDelayTimer()
	c.UpdateSoundTimer()

	for i := 0; i < len(c.KeyState); i++ {
		if c.KeyState[i] == 1 {
			fmt.Println("La touche qui a pour valeur : ", i, " est enfoncée")
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
	return 320, 160
}

func OpenWindowEbiten(cpu *CPU) {
	ebiten.SetWindowSize(1280, 640)
	ebiten.SetWindowTitle("Chip 8")
	if err := ebiten.RunGame(cpu); err != nil {
		log.Fatal(err)
	}
}

func (c *CPU) KeyPress() {
	if ebiten.IsKeyPressed(ebiten.Key0) || ebiten.IsKeyPressed(ebiten.KeyNumpad0) {
		c.KeyState[0] = 1
	} else {
		c.KeyState[0] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key1) || ebiten.IsKeyPressed(ebiten.KeyNumpad1) {
		c.KeyState[1] = 1
	} else {
		c.KeyState[1] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key2) || ebiten.IsKeyPressed(ebiten.KeyNumpad2) {
		c.KeyState[2] = 1
	} else {
		c.KeyState[2] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key3) || ebiten.IsKeyPressed(ebiten.KeyNumpad3) {
		c.KeyState[3] = 1
	} else {
		c.KeyState[3] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key4) || ebiten.IsKeyPressed(ebiten.KeyNumpad4) {
		c.KeyState[4] = 1
	} else {
		c.KeyState[4] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key5) || ebiten.IsKeyPressed(ebiten.KeyNumpad5) {
		c.KeyState[5] = 1
	} else {
		c.KeyState[5] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key6) || ebiten.IsKeyPressed(ebiten.KeyNumpad6) {
		c.KeyState[6] = 1
	} else {
		c.KeyState[6] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key7) || ebiten.IsKeyPressed(ebiten.KeyNumpad7) {
		c.KeyState[7] = 1
	} else {
		c.KeyState[7] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key8) || ebiten.IsKeyPressed(ebiten.KeyNumpad8) {
		c.KeyState[8] = 1
	} else {
		c.KeyState[8] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key9) || ebiten.IsKeyPressed(ebiten.KeyNumpad9) {
		c.KeyState[9] = 1
	} else {
		c.KeyState[9] = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		c.KeyState[0xA] = 1
	} else {
		c.KeyState[0xA] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		c.KeyState[0xB] = 1
	} else {
		c.KeyState[0xB] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		c.KeyState[0xC] = 1
	} else {
		c.KeyState[0xC] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.KeyState[0xD] = 1
	} else {
		c.KeyState[0xD] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		c.KeyState[0xE] = 1
	} else {
		c.KeyState[0xE] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		c.KeyState[0xF] = 1
	} else {
		c.KeyState[0xF] = 0
	}
}

func (c *CPU) UpdateDelayTimer() {
	if c.DT > 0 {
		c.DT--
	}
}

func (c *CPU) UpdateSoundTimer() {
	if c.ST > 0 {
		c.ST--
	}
}

func (c *CPU) SoundEnabled() bool {
	return c.ST > 0
}
