package reg

import "sap-1/bus"

/*
   Reg represents an 8-bit register.

                        ▲
                   OutBus (8)
                ┏━━━━━━━┷━━━━━━━━┓
    InBus (8) ▶ ┃                ┠───── L̅x
                ┃    Register   ◁┠───── CLK
                ┃    ▢▢▢▢▢▢▢▢    ┃
                ┗━━━━━━━━━━━━━━━━┛
*/
type Reg struct {
	InBus  *bus.Bus // Data (In) Bus
	OutBus *bus.Bus // Data (Out) Bus

	NLx func() bool // Load State

	Data uint8
}

// NewReg: Constructs a new Register.
func NewReg(InBus, OutBus *bus.Bus, NLx func() bool) *Reg {
	return &Reg{InBus: InBus, OutBus: OutBus, NLx: NLx}
}

/*
	Clk: According to the Load State,
	- Reads the bus data into the register.
	- Writes the register's data to the output bus.
*/
func (reg *Reg) Clk() {
	if !reg.NLx() {
		reg.Data = reg.InBus.Data
		reg.OutBus.Write(reg.Data)
	}
}
