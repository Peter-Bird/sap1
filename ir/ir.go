package ir

import "sap-1/bus"

/*
   IR simulates an 8-bit instruction register

             ┏━━━━━━━━━━━━━━━━━┓
    NLi ─────┨                 ┃◀──  Wbus (8)
    CLK ─────┨▷      IR        ┃
    CLR ─────┨    ▢▢▢▢ ▢▢▢▢    ┃
    NEi ─────┨                 ┠──── Wbus (4) ▶
             ┗━━━━━━━━┳━━━━━━━━┛
                   Ibus (4)
                     ▼
*/
type IR struct {
	Wbus *bus.Bus
	Ibus *bus.Bus

	NLi func() bool
	NEi func() bool

	Data uint8
}

// NewIR: Creates a new instance of an Instruction Register (IR).
func NewIR(Wbus, Ibus *bus.Bus, NLi, NEi func() bool) *IR {
	return &IR{Wbus: Wbus, Ibus: Ibus, NLi: NLi, NEi: NEi}
}

// Clr: Sets the Instruction Register to zero.
func (ir *IR) Clr() { ir.Data = 0 }

/*
	Clk: According to the control lines
	- NLi: Loads Data from the bus in to the Instruction Register.
		   Writes the instrcution (Left nibble) to the Controller's bus.
	- NEi: Sends the Address in the instruction (Right nibble) to the bus.
*/
func (ir *IR) Clk() {
	if !ir.NLi() {
		ir.Data = ir.Wbus.Data
		ir.Ibus.Write(ir.Data & 0xF0 >> 4) // Instruction
	}
	if !ir.NEi() {
		ir.Wbus.Write(ir.Data & 0x0F) // Address
	}
}
