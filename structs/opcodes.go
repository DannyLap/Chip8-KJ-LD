package structs

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

func (c *CPU) OpcodesReading(data [2]byte) { //il faudra utiliser les opcodes par la suite
	opcode := uint16(data[0])<<8 | uint16(data[1])
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode {
		case 0x00E0:
		case 0x00EE:
		}
	case 0x1000:
		nn := byte(opcode & 0x0FFF)
		c.PC = nn
	case 0x2000:
		//nnn := opcode & 0x0FFF
		//Call subroutine at nnn. The interpreter increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
		//z'ai po compris ça
	case 0x3000:
		vX := opcode & 0x0F00
		kk := byte(opcode & 0x00FF)
		if c.Registers[vX] == kk {
			c.PC += 2
		}
	case 0x4000:
		vX := opcode & 0x0F00
		kk := byte(opcode & 0x00FF)
		if c.Registers[vX] == kk {
			c.PC += 2
		}
	case 0x5000:
		vX := opcode & 0x0F00
		vY := opcode & 0x00F0
		if c.Registers[vX] == c.Registers[vY] {
			c.PC += 2
		}
	case 0x6000:
		vX := opcode & 0x0F00
		kk := byte(opcode & 0x00FF)
		c.Registers[vX] = kk
		// Set Vx = kk. The interpreter puts the value kk into register Vx.
	case 0x7000:
		vX := opcode & 0x0F00
		kk := byte(opcode & 0x00FF)
		c.Registers[vX] = kk
		//Set Vx = Vx + kk. Adds the value kk to the value of register Vx, then stores the result in Vx.
	case 0x8000:
		switch opcode & 0xF00F {
		case 0x8000:
			vX := opcode & 0x0F00
			vY := opcode & 0x00F0
			c.Registers[vX] = c.Registers[vY]
			//Set Vx = Vy. Stores the value of register Vy in register Vx
		case 0x8001:
			//Set Vx = Vx OR Vy. Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A
			// bitwise OR compares the corresponding bits from two values, and if either bit is 1, then the same bit in the
			// result is also 1. Otherwise, it is 0.
			//trop compliqué pour le moment
		case 0x8002:
			// Set Vx = Vx AND Vy. Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx.
			// A bitwise AND compares the corresponding bits from two values, and if both bits are 1, then the same bit
			// in the result is also 1. Otherwise, it is 0.
			//trop compliqué pour le moment
		case 0x8003:
			// Set Vx = Vx XOR Vy. Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result
			// in Vx. An exclusive OR compares the corresponding bits from two values, and if the bits are not both the
			// same, then the corresponding bit in the result is set to 1. Otherwise, it is 0
		case 0x8004:
			//Set Vx = Vx + Vy, set VF = carry. The values of Vx and Vy are added together. If the result is greater
			// than 8 bits (i.e., ¿ 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored
			// in Vx
		case 0x8005:
			// Set Vx = Vx - Vy, set VF = NOT borrow. If Vx ¿ Vy, then VF is set to 1, otherwise 0. Then Vy is
			// subtracted from Vx, and the results stored in Vx
		case 0x8006:
			// Set Vx = Vx SHR 1. If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is
			// divided by 2
		case 0x8007:
			// 	Set Vx = Vy - Vx, set VF = NOT borrow. If Vy ¿ Vx, then VF is set to 1, otherwise 0. Then Vx is
			// subtracted from Vy, and the results stored in Vx.
		case 0x800E:
			// Set Vx = Vx SHL 1. If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is
			// multiplied by 2.

		}
	case 0xA000:
		nnn := byte(opcode & 0x0FFF)
		c.Registers[nnn] = c.Registers[c.I]
		c.I = nnn
		//Set I = nnn. The value of register I is set to nnn.
	case 0xB000:
		nnn := byte(opcode & 0x0FFF)
		c.PC = nnn + c.Registers[0]
		// Jump to location nnn + V0. The program counter is set to nnn plus the value of V0
	case 0xC000:
		//Set Vx = random byte AND kk. The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
	case 0xD000:
		// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision. The interpreter reads n
		// bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen
		// at coordinates (Vx, Vy). Sprites are XOR’d onto the existing screen. If this causes any pixels to be erased,
		// VF is set to 1, otherwise it is set to 0. If the sprite is positioned so part of it is outside the coordinates of
		// the display, it wraps around to the opposite side of the screen.
	case 0xE000:
		switch opcode & 0xF0FF {
		case 0xE09E:
			//Skip next instruction if key with the value of Vx is pressed. Checks the keyboard, and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.
		case 0xE0A1:
			//Skip next instruction if key with the value of Vx is not pressed. Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2
		}
	case 0xF000:
		switch opcode & 0xF0FF {
		case 0xF007:
			//Set Vx = delay timer value. The value of DT is placed into Vx.
		case 0xF00A:
			//Wait for a key press, store the value of the key in Vx. All execution stops until a key is pressed, then the value of that key is stored in Vx.
		case 0xF015:
			//Set delay timer = Vx. Delay Timer is set equal to the value of Vx.
		case 0xF018:
			//Set sound timer = Vx. Sound Timer is set equal to the value of Vx
		case 0xF01E:
			//Set I = I + Vx. The values of I and Vx are added, and the results are stored in I.
		case 0xF029:
			//Set I = location of sprite for digit Vx. The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal font. To obtain this value, multiply VX by 5 (all font data stored in first 80 bytes of memory)
		case 0xF033:
			//Store BCD representation of Vx in memory locations I, I+1, and I+2. The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.
		case 0xF055:
			//Stores V0 to VX in memory starting at address I. I is then set to I + x + 1
		case 0xF065:
			//Fills V0 to VX with values from memory starting at address I. I is then set to I + x + 1.
		}
	}
}
