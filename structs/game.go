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

type Input struct {
	KeyState        [16]byte
	WaitingForInput bool
}

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

	for i := 0; i < len(c.Input.KeyState); i++ {
		if c.Input.KeyState[i] == 1 {
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
	keyPress := [16]bool{
		ebiten.IsKeyPressed(ebiten.Key0) || ebiten.IsKeyPressed(ebiten.KeyNumpad0),
		ebiten.IsKeyPressed(ebiten.Key1) || ebiten.IsKeyPressed(ebiten.KeyNumpad1),
		ebiten.IsKeyPressed(ebiten.Key2) || ebiten.IsKeyPressed(ebiten.KeyNumpad2),
		ebiten.IsKeyPressed(ebiten.Key3) || ebiten.IsKeyPressed(ebiten.KeyNumpad3),
		ebiten.IsKeyPressed(ebiten.Key4) || ebiten.IsKeyPressed(ebiten.KeyNumpad4),
		ebiten.IsKeyPressed(ebiten.Key5) || ebiten.IsKeyPressed(ebiten.KeyNumpad5),
		ebiten.IsKeyPressed(ebiten.Key6) || ebiten.IsKeyPressed(ebiten.KeyNumpad6),
		ebiten.IsKeyPressed(ebiten.Key7) || ebiten.IsKeyPressed(ebiten.KeyNumpad7),
		ebiten.IsKeyPressed(ebiten.Key8) || ebiten.IsKeyPressed(ebiten.KeyNumpad8),
		ebiten.IsKeyPressed(ebiten.Key9) || ebiten.IsKeyPressed(ebiten.KeyNumpad9),
		ebiten.IsKeyPressed(ebiten.KeyQ),
		ebiten.IsKeyPressed(ebiten.KeyB),
		ebiten.IsKeyPressed(ebiten.KeyC),
		ebiten.IsKeyPressed(ebiten.KeyD),
		ebiten.IsKeyPressed(ebiten.KeyE),
		ebiten.IsKeyPressed(ebiten.KeyF)}
	for keyIndex := range keyPress {
		if keyPress[keyIndex] {
			c.Input.KeyState[keyIndex] = 1
		} else {
			c.Input.KeyState[keyIndex] = 0
		}
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

func (c *CPU) CatchInput(keyIndex byte) {
	c.Registers[c.waitingForInputRegisterInput] = keyIndex
	c.WaitingForInput = false
}
