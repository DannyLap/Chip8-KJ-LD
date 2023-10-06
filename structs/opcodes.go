package structs

import (
	"fmt"
	"math/rand"
	"time"
)

func (c *CPU) AddOpcodesToCPU() {
	c.Opcodes = []byte{
		0x00, 0xE0, 0x00, 0xEE, 0x00, 0xC0, 0x00, 0xFB,
		0x00, 0xFC, 0x00, 0xFD, 0x00, 0xFE, 0x00, 0xFF,
		0x10, 0x00, 0x20, 0x00, 0x30, 0x00, 0x40, 0x00,
		0x50, 0x00, 0x60, 0x00, 0x70, 0x00, 0x80, 0x00,
		0x80, 0x01, 0x80, 0x02, 0x80, 0x03, 0x80, 0x04,
		0x80, 0x05, 0x80, 0x06, 0x80, 0x07, 0x80, 0x0E,
		0x90, 0x00, 0xA0, 0x00, 0xB0, 0x00, 0xC0, 0x00,
		0xD0, 0x00, 0xE0, 0x9E, 0xE0, 0xA1, 0xF0, 0x07,
		0xF0, 0x0A, 0xF0, 0x15, 0xF0, 0x18, 0xF0, 0x1E,
		0xF0, 0x29, 0xF0, 0x33, 0xF0, 0x55, 0xF0, 0x65,
	}
}

func (c *CPU) OpcodesReading() {
	opcode := uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1])
	fmt.Print(c.PC)
	fmt.Printf(": 0x%X %X\n", c.Memory[c.PC], c.Memory[c.PC+1])
	fmt.Println(c.PC, c.Memory[c.PC], c.Memory[c.PC+1])
	fmt.Println(c.Registers)
	c.PC += 2
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode {
		case 0x00E0:
			c.ClearScreen()
			// Clear the display
		case 0x00EE:
			fmt.Println("HEY !!!!!!!!!!!!!!!!!!!!!!!")
			c.PC = uint16(c.Stack[c.SP])
			c.SP--
			// Return from a subroutine.The interpreter sets the program counter to the address at the top of the stack,
			// then subtracts 1 from the stack pointer
		}
	case 0x1000:
		nnn := opcode & 0x0FFF
		c.PC = nnn
	case 0x2000:
		fmt.Println("ON UTILISE CA TA MERE")
		nnn := opcode & 0x0FFF
		c.SP++
		c.Stack[c.SP] = c.PC
		c.PC = nnn
		//Call subroutine at nnn. The interpreter increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
	case 0x3000:
		x := (opcode & 0x0F00) / 256
		nn := byte(opcode & 0x00FF)
		if x == 9 {
			fmt.Println("V(", x, ") = ", c.Registers[x], ", nn =", nn, "ils sont diff ", c.Registers[x] != nn)
		}
		if c.Registers[x] == nn {
			c.PC += 2
		}
	case 0x4000:
		x := (opcode & 0x0F00) / 256
		nn := byte(opcode & 0x00FF)
		if x == 9 {
			fmt.Println("V(", x, ") = ", c.Registers[x], "nn =", nn, "ils sont égaux ", c.Registers[x] == nn)
		}
		if c.Registers[x] != nn {
			c.PC += 2
		}
	case 0x5000:
		x := (opcode & 0x0F00) / 256
		y := (opcode & 0x00F0) / 16
		if x == 9 || y == 9 {
			fmt.Println("V(", x, ") = ", c.Registers[x], "et V(", y, ") =", c.Registers[y], "ils sont égaux ", c.Registers[x] == c.Registers[y])
		}
		if c.Registers[x] == c.Registers[y] {
			c.PC += 2
		}
	case 0x6000:
		x := (opcode & 0x0F00) / 256
		kk := byte(opcode & 0x00FF)
		if x == 9 {
			fmt.Println("v(", x, ") = ", c.Registers[x], "et kk =", kk)
		}
		c.Registers[x] = kk
		// Set Vx = kk. The interpreter puts the value kk into register Vx.
	case 0x7000:
		x := (opcode & 0x0F00) / 256
		kk := byte(opcode & 0x00FF)
		c.Registers[x] += kk
		if x == 9 {
			fmt.Println("v(", x, ") = ", c.Registers[x])
		}
		//Set Vx = Vx + kk. Adds the value kk to the value of register Vx, then stores the result in Vx.
	case 0x8000:
		switch opcode & 0xF00F {
		case 0x8000:
			x := (opcode & 0x0F00) / 256
			y := (opcode & 0x00F0) / 16
			c.Registers[x] = c.Registers[y]
			if x == 9 || y == 9 {
				fmt.Println("v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			//Set Vx = Vy. Stores the value of register Vy in register Vx
		case 0x8001:
			x := (opcode & 0x0F00) / 256
			y := (opcode & 0x00F0) / 16
			if x == 9 || y == 9 {
				fmt.Println("avant v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			c.Registers[x] = c.Registers[x] | c.Registers[y]
			if x == 9 || y == 9 {
				fmt.Println(" donc v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			//Set Vx = Vx OR Vy. Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A
			// bitwise OR compares the corresponding bits from two values, and if either bit is 1, then the same bit in the
			// result is also 1. Otherwise, it is 0.
		case 0x8002:
			x := (opcode & 0x0F00) / 256
			y := (opcode & 0x00F0) / 16
			if x == 9 || y == 9 {
				fmt.Println("avant v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			c.Registers[x] = c.Registers[x] & c.Registers[y]
			if x == 9 || y == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			// Set Vx = Vx AND Vy. Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx.
			// A bitwise AND compares the corresponding bits from two values, and if both bits are 1, then the same bit
			// in the result is also 1. Otherwise, it is 0.
			//trop compliqué pour le moment
		case 0x8003:
			x := (opcode & 0x0F00) / 256
			y := (opcode & 0x00F0) / 16
			if x == 9 || y == 9 {
				fmt.Println("avant v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			c.Registers[x] = c.Registers[x] ^ c.Registers[y]
			if x == 9 || y == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			// Set Vx = Vx XOR Vy. Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result
			// in Vx. An exclusive OR compares the corresponding bits from two values, and if the bits are not both the
			// same, then the corresponding bit in the result is set to 1. Otherwise, it is 0
		case 0x8004:
			x := (opcode & 0x0F00) / 256
			y := (opcode & 0x00F0) / 16
			if x == 9 || y == 9 {
				fmt.Println("avant v(", x, ") = ", uint16(c.Registers[x]), "v(", y, ") = ", uint16(c.Registers[y]))
				fmt.Println("v(", x, ") + v(", y, ") = ", uint16(c.Registers[x])+uint16(c.Registers[y]))
			}
			if uint16(c.Registers[x])+uint16(c.Registers[y]) > 255 {
				c.Registers[x] += c.Registers[y]
				c.Registers[0xF] = 1
			} else {
				c.Registers[x] += c.Registers[y]
				c.Registers[0xF] = 0
			}
			if x == 9 || y == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			//Set Vx = Vx + Vy, set VF = carry. The values of Vx and Vy are added together. If the result is greater
			// than 8 bits (i.e., ¿ 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored
			// in Vx
		case 0x8005:
			x := (opcode & 0x0F00) / 256
			y := (opcode & 0x00F0) / 16
			if x == 9 || y == 9 {
				fmt.Println("avant v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			if c.Registers[x] >= c.Registers[y] {
				c.Registers[x] -= c.Registers[y]
				c.Registers[0xF] = 1
			} else {
				c.Registers[x] -= c.Registers[y]
				c.Registers[0xF] = 0
			}
			if x == 9 || y == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			// Set Vx = Vx - Vy, set VF = NOT borrow. If Vx ¿ Vy, then VF is set to 1, otherwise 0. Then Vy is
			// subtracted from Vx, and the results stored in Vx
		case 0x8006:
			x := (opcode & 0x0F00) / 256
			if x == 9 {
				fmt.Println("avant v(", x, ") = ", c.Registers[x])
			}
			if c.Registers[x]&1 == 1 {
				c.Registers[x] >>= 1
				c.Registers[0xF] = 1
			} else {
				c.Registers[x] >>= 1
				c.Registers[0xF] = 0
			}
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			// Set Vx = Vx SHR 1. If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is
			// divided by 2
		case 0x8007:
			x := (opcode & 0x0F00) / 256
			y := (opcode & 0x00F0) / 16
			if x == 9 || y == 9 {
				fmt.Println("avant v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			if c.Registers[y] >= c.Registers[x] {
				c.Registers[x] = c.Registers[y] - c.Registers[x]
				c.Registers[0xF] = 1
			} else {
				c.Registers[x] = c.Registers[y] - c.Registers[x]
				c.Registers[0xF] = 0
			}
			if x == 9 || y == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y])
			}
			// 	Set Vx = Vy - Vx, set VF = NOT borrow. If Vy ¿ Vx, then VF is set to 1, otherwise 0. Then Vx is
			// subtracted from Vy, and the results stored in Vx.
		case 0x800E:
			x := (opcode & 0x0F00) / 256
			if x == 9 {
				fmt.Println("avant v(", x, ") = ", c.Registers[x])
			}
			if c.Registers[x]&0x80 == 0x80 {
				c.Registers[x] <<= 1
				c.Registers[0xF] = 1
			} else {
				c.Registers[x] <<= 1
				c.Registers[0xF] = 0
			}
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			// Set Vx = Vx SHL 1. If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is
			// multiplied by 2.
		}
	case 0x9000:
		x := (opcode & 0x0F00) / 256
		y := (opcode & 0x00F0) / 16
		// if x == 9 || y == 9 {
		// 	fmt.Println("v(", x, ") = ", c.Registers[x], "v(", y, ") = ", c.Registers[y], "ils sont diff", c.Registers[x] != c.Registers[y])
		// }
		if c.Registers[x] != c.Registers[y] {
			c.PC += 2
		}
		//Skip next instruction if Vx != Vy. The values of Vx and Vy are compared, and if they are not equal, the program counter is increased by 2.
	case 0xA000:
		nnn := opcode & 0x0FFF
		c.I = nnn
		fmt.Println("I = ", c.I)
		//Set I = nnn. The value of register I is set to nnn.
	case 0xB000:
		nnn := opcode & 0x0FFF
		c.PC = nnn + uint16(c.Registers[0])
		// Jump to location nnn + V0. The program counter is set to nnn plus the value of V0
	case 0xC000:
		rand.Seed(time.Now().UnixNano())
		x := (opcode & 0x0F00) / 256
		kk := byte(opcode & 0x00FF)
		rdByte := byte(rand.Intn(256))
		c.Registers[x] = rdByte & kk
		if x == 9 {
			fmt.Println("donc v(", x, ") = ", c.Registers[x])
		}
		//Set Vx = random byte AND kk. The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
	case 0xD000:
		x := (opcode & 0x0F00) / 256
		y := (opcode & 0x00F0) / 16
		n := byte(opcode & 0x000F)
		collision := false
		fmt.Println(c.Memory[c.I : c.I+uint16(n)])
		for k := byte(0); k < n; k++ {
			spriteByte := c.Memory[c.I+uint16(k)] // Lire l'octet du sprite depuis la mémoire à l'emplacement I

			for l := byte(0); l < 8; l++ {
				// Lire le bit du sprite à partir du MSB (bit le plus significatif)
				spriteBit := (spriteByte >> (7 - l)) & 0x1
				i := (c.Registers[x] + l) % 64 // Gestion du wrapping horizontal
				j := (c.Registers[y] + k) % 32 // Gestion du wrapping vertical
				oldPixel := c.Screen[i][j]

				// XOR entre le pixel actuel et le bit du sprite
				c.Screen[i][j] ^= spriteBit

				// Si le pixel était allumé avant et est éteint maintenant, il y a collision
				if oldPixel == 1 && c.Screen[i][j] == 0 {
					collision = true
				}
			}
		}

		// Mettre à jour le drapeau VF en fonction de la collision
		if collision {
			c.Registers[0xF] = 1
		} else {
			c.Registers[0xF] = 0
		}
		// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision. The interpreter reads n
		// bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen
		// at coordinates (Vx, Vy). Sprites are XOR’d onto the existing screen. If this causes any pixels to be erased,
		// VF is set to 1, otherwise it is set to 0. If the sprite is positioned so part of it is outside the coordinates of
		// the display, it wraps around to the opposite side of the screen.
	case 0xE000:
		switch opcode & 0xF0FF {
		case 0xE09E:
			// todo
			//  Skip next instruction if key with the value of Vx is pressed. Checks the keyboard,
			//  and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.

			x := int16((opcode & 0x0F00) / 256)

			fmt.Println("down x = ", c.Registers[x])
			fmt.Println(c.KeyState)

			if c.KeyState[c.Registers[x]] == 1 {
				c.PC += 2
			}

		case 0xE0A1:
			// todo
			//  Skip next instruction if key with the value of Vx is not pressed. Checks the keyboard,
			//  and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2

			x := int16((opcode & 0x0F00) / 256)
			//key := StringToHexa(c.Key)

			fmt.Println("up x = ", c.Registers[x])
			fmt.Println(c.KeyState)

			fmt.Println("op 2 deeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
			if c.KeyState[c.Registers[x]] == 0 {
				c.PC += 2
			}
		}
	case 0xF000:
		switch opcode & 0xF0FF {
		case 0xF007:
			x := (opcode & 0x0F00) / 256
			c.Registers[x] = byte(c.DT)
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			//Set Vx = delay timer value. The value of DT is placed into Vx.
		case 0xF00A:
			x := (opcode & 0x0F00) / 256
			y := byte(0)
			fmt.Println(c.KeyState)

			for c.KeyState[y] != 1 {
				if int(y) == len(c.KeyState)-1 {
					y = 0
				} else {
					y++
				}
			}
			c.Registers[x] = y
			fmt.Println("CACACACACACACACACACACACACACACACACACACACACACACACACACACACACACACA")
			//Wait for a key press, store the value of the key in Vx. All execution stops until a key is pressed,
			//then the value of that key is stored in Vx.
		case 0xF015:
			x := (opcode & 0x0F00) / 256
			c.DT = uint16(c.Registers[x])
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			//Set delay timer = Vx. Delay Timer is set equal to the value of Vx.
		case 0xF018:
			x := (opcode & 0x0F00) / 256
			c.ST = uint16(c.Registers[x])
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			//Set sound timer = Vx. Sound Timer is set equal to the value of Vx
		case 0xF01E:
			x := (opcode & 0x0F00) / 256
			c.I += uint16(c.Registers[x])
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			//Set I = I + Vx. The values of I and Vx are added, and the results are stored in I.
		case 0xF029:
			x := (opcode & 0x0F00) / 256
			c.I = uint16(c.Registers[x]) * 5 // + 0x200 ?
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			//Set I = location of sprite for digit Vx. The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal font. To obtain this value, multiply VX by 5 (all font data stored in first 80 bytes of memory)
		case 0xF033:
			x := (opcode & 0x0F00) / 256
			Vx := c.Registers[x]
			c.Memory[c.I+2] = Vx % 10
			Vx /= 10
			c.Memory[c.I+1] = Vx % 10
			Vx /= 10
			c.Memory[c.I] = Vx
			if x == 9 {
				fmt.Println("donc v(", x, ") = ", c.Registers[x])
			}
			//Store BCD representation of Vx in memory locations I, I+1, and I+2. The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.
		case 0xF055:
			x := (opcode & 0x0F00) / 256
			for k := uint16(0); k <= x; k++ {
				c.Memory[c.I+k] = c.Registers[k]
			}
			c.I += x + 1

			//Stores V0 to VX in memory starting at address I. I is then set to I + x + 1
		case 0xF065:
			x := (opcode & 0x0F00) / 256
			for k := uint16(0); k <= x; k++ {
				fmt.Println("k = ", k, " V(k) = ", c.Registers[k], " I+k = ", c.I+k, " Memory(I+k)=", c.Memory[c.I+k])
				c.Registers[k] = c.Memory[c.I+k]
			}
			c.I += x + 1

			//Fills V0 to VX with values from memory starting at address I. I is then set to I + x + 1.
		}
	}
}
