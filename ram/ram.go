package ram

import "sap-1/bus"

/*
   RAM represents a 16x8 bit Random Access Memory.

               Rbus (4)
                  ▼
            ┏━━━━━┷━━━━━━┓
            ┃    RAM     ┃
    NCE ────┨  ▢▢▢▢▢▢▢▢  ┠─ Wbus (8) ▶
            ┃  ▢▢▢▢▢▢▢▢  ┃
            ┃    ...     ┃
            ┃  ▢▢▢▢▢▢▢▢  ┃
            ┗━━━━━━━━━━━━┛
*/
type Ram struct {
	Wbus *bus.Bus
	Rbus *bus.Bus

	NCE func() bool

	Data [16]uint8
}

//	NewRam: Constructs a new instance of Ram
func NewRam(Wbus, Rbus *bus.Bus, NCE func() bool) *Ram {
	return &Ram{Wbus: Wbus, Rbus: Rbus, NCE: NCE}
}

//	Exec: Writes the data in Ram[address] to the bus
func (ram *Ram) Exec() {
	if !ram.NCE() {
		ram.Wbus.Write(ram.Data[ram.Rbus.Data])
	}
}

//	Write: Writes a byte of Data to an address in Ram.
func (m *Ram) Write(addr uint8, Data uint8) {
	m.Data[addr] = Data
}

//	BurnRam: Writes a program into Ram
func (ram *Ram) BurnRam() {
	for i, val := range Program {
		ram.Write(uint8(i), val)
	}
}
