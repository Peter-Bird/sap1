package alu

import "sap-1/bus"

/*
   ALU represents an 8-bit Arthimetic Logic Unit.

                  Abus (8)
                      ▼
              ┏━━━━━━━┷━━━━━━━━┓
              ┃                ┠───── Su
              ┃      ALU       ┃
   ◀ Wbus (8) ┃    ▢▢▢▢▢▢▢▢    ┠───── Eu
              ┗━━━━━━━┳━━━━━━━━┛
                      ▲
                  Bbus (8)
*/
type ALU struct {
	Abus *bus.Bus
	Bbus *bus.Bus
	Wbus *bus.Bus

	Su func() bool
	Eu func() bool

	Data uint8
}

//	NewALU: Constructs a new instance of an ALU
func NewALU(Abus, Bbus, Wbus *bus.Bus, Su, Eu func() bool) *ALU {
	return &ALU{Abus: Abus, Bbus: Bbus, Wbus: Wbus, Su: Su, Eu: Eu}
}

/*
	Exec: Performs addition or sustraction of the ACC and Breg data
	- Eu: Write the data to the bus
	- SU: If set Subtract otherwise Add
*/
func (a *ALU) Exec() {
	if a.Su() {
		a.Data = a.Abus.Data - a.Bbus.Data
	} else {
		a.Data = a.Abus.Data + a.Bbus.Data
	}

	if a.Eu() {
		a.Wbus.Write(a.Data)
	}
}
