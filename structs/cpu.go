package structs

type CPU struct {
	Registers []byte
	PC        byte
	I         byte
	SP        byte
	DT        byte
	ST        byte
	Memory    [4096]byte
}

func (c *CPU) AddROMToMemory(data []byte) {
	for i, b := range data {
		c.Memory[0x200+i] = b
	}
}
