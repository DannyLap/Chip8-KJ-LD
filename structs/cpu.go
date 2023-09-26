package structs

type CPU struct {
	Registers []byte
	PC        byte
	I         byte
	SP        byte
	DT        byte
	ST        byte
	Memory    []byte
}
