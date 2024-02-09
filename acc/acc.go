package acc

import "sap-1/bus"

/*
   Acc simulates an 8-bit accumulator.

              ┏━━━━━━━━━━━━━━━━┓
   Wbus (8) ▶ ┃                ┠───── NLa
              ┃  Accumulator  ◁┠───── CLK
   ◀ Wbus (8) ┃   ▢▢▢▢ ▢▢▢▢    ┠───── Ea
              ┗━━━━━━━┳━━━━━━━━┛
                   Abus (8)
                      ▼
*/
type Acc struct {
	Wbus *bus.Bus
	Abus *bus.Bus

	NLa func() bool
	Ea  func() bool

	Data uint8
}

//	NewAcc: Constructs an instance of an accumulator.
func NewAcc(Wbus, Abus *bus.Bus, NLa, Ea func() bool) *Acc {
	return &Acc{Wbus: Wbus, Abus: Abus, NLa: NLa, Ea: Ea}
}

/*
	Clk: According to the control lines
	- Ea: Writes the accumulator's data to the bus
	- NLa: Reads the data from the bus into data
	       Writes the data to the ALU's bus.
*/
func (acc *Acc) Clk() {

	if acc.Ea() {
		acc.Wbus.Write(acc.Data)
	}
	if !acc.NLa() {
		acc.Data = acc.Wbus.Data
		acc.Abus.Write(acc.Data)
	}
}
