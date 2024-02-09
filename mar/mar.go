package mar

import "sap-1/bus"

/*
   MAR represents a 4-bit Memory Address Register.

             ┏━━━━━━━━━━━━┓
    NLm ─────┨    MAR     ┃◀──  Wbus (4)
    CLK ─────┨▷   ▢▢▢▢    ┃
             ┗━━━━━┳━━━━━━┛
                 Rbus (4)
                   ▼
*/
type MAR struct {
	Wbus *bus.Bus
	Rbus *bus.Bus

	NLm func() bool

	Data uint8
}

// NewMAR: Creates a new 4-bit Memory Address Register (MAR).
func NewMAR(Wbus *bus.Bus, Rbus *bus.Bus, NLm func() bool) *MAR {
	return &MAR{Wbus: Wbus, Rbus: Rbus, NLm: NLm}
}

// Clk: Sets the value of the memory address register.
func (mar *MAR) Clk() {
	if !mar.NLm() {
		mar.Data = mar.Wbus.Data & 0x0F // Right 4 bits
		mar.Rbus.Write(mar.Data)
	}
}
