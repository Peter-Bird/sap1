package pc

import "sap-1/bus"

/*
   PC represents a 4-bit Program Counter.

              ┏━━━━━━━━━━━━┓
      Cp ─────┨            ┃
    NCLK ─────┨▷    PC     ┠──── Wbus (4) ▶
    NCLR ─────┨    ▢▢▢▢    ┃
      Ep ─────┨            ┃
              ┗━━━━━━━━━━━━┛
*/
type PC struct {
	Wbus *bus.Bus

	Cp func() bool
	Ep func() bool

	Data uint8
}

// NewPC: Creates a new instance of a Program Counter (PC).
func NewPC(Wbus *bus.Bus, Cp, Ep func() bool) *PC {
	return &PC{Wbus: Wbus, Cp: Cp, Ep: Ep}
}

// NClr: Sets the Program Counter to zero.
func (pc *PC) NClr() { pc.Data = 0 }

/*
	NClk: On the clock's dropping edge
	- Ep: Writes the PC's value to the bus
	- Cp: Increment the PC
*/
func (pc *PC) NClk() {
	if pc.Ep() {
		pc.Wbus.Write(pc.Data & 0x0F)
	}
	if pc.Cp() {
		pc.Data = (pc.Data + 1) & 0x0F
	}
}
