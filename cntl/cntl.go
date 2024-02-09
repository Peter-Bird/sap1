package cntl

import (
	"sap-1/bus"
	"sap-1/clk"
	"sap-1/cntlBus"
)

/*
   Cntl represents an 12-bit Controller Sequencer.

          Ibus (4)
            ▼
    ┏━━━━━━━┷━━━━━━━━┓
    ┃                ┠──── CLK
    ┃   Controller   ┠──── NCLK
    ┃  ▢▢▢▢▢▢▢▢▢▢▢▢  ┠──── CLR
    ┃                ┠──── NCLR
    ┗━━━━━━━┳━━━━━━━━┛
        CntlBus (12)
            ▼
*/
type Cntl struct {
	Clock *clk.Clock

	Ibus *bus.Bus
	Cbus *cntlBus.CntlBus

	onClk  []func()
	onNClk []func()
	onClr  []func()
	onNClr []func()

	MicroProg []uint16
	MicroStep int

	Data uint16
}

/*
	NewCntl: Constructs a new instance of a Controller.
	- Sets Fetch as the initial Micro Program
	- Sets Fetch[0] as the initial Micro Code to execute
*/
func NewCntl(Clk *clk.Clock, Ibus *bus.Bus, Cbus *cntlBus.CntlBus) *Cntl {
	return &Cntl{
		Clock: Clk, Ibus: Ibus, Cbus: Cbus, MicroProg: Fetch, Data: Fetch[0],
	}
}

// Exec: Loads the Micro Program into the controller.
func (cntl *Cntl) Exec() { cntl.LoadMicroProg() }

// Clk: Triggers the registered Component's Clock event.
func (cntl *Cntl) Clk() { cntl.Trig(cntl.onClk) }

/*
	NClk: Loads the next Micro Code and then
	Triggers the registered Component's (Not) Clock event.
*/
func (cntl *Cntl) NClk() {
	cntl.NextMicro()
	cntl.Trig(cntl.onNClk)
}

// Start: Starts the run clearing registered components.
func (cntl *Cntl) Start() {
	cntl.Trig(cntl.onNClr)
	cntl.Trig(cntl.onClr)
}
