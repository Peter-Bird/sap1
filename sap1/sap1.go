package sap1

import (
	"sap-1/acc"
	"sap-1/alu"
	"sap-1/bus"
	"sap-1/clk"
	"sap-1/cntl"
	"sap-1/cntlBus"
	"sap-1/dbg"

	"sap-1/dis"
	"sap-1/ir"
	"sap-1/mar"
	"sap-1/pc"
	"sap-1/ram"
	"sap-1/reg"
)

type Sap1 struct {
	Clk *clk.Clock

	Acc  *acc.Acc
	Alu  *alu.ALU
	BReg *reg.Reg
	OReg *reg.Reg
	Dis  *dis.Dis

	Pc  *pc.PC
	Mar *mar.MAR
	Ram *ram.Ram
	Ir  *ir.IR

	Cntl *cntl.Cntl

	ABus *bus.Bus
	BBus *bus.Bus
	DBus *bus.Bus
	OBus *bus.Bus
	RBus *bus.Bus
	IBus *bus.Bus
	CBus *cntlBus.CntlBus

	Dbg *dbg.Dbg
}

func NewSap1() *Sap1 {

	// Create Buses
	abus := bus.NewBus()
	bbus := bus.NewBus()
	dbus := bus.NewBus()
	obus := bus.NewBus()
	rbus := bus.NewBus()
	ibus := bus.NewBus()
	cbus := cntlBus.NewCntlBus()

	clk := clk.NewClock(100)

	// Create Components
	pc := pc.New(dbus, cbus.Cp, cbus.Ep)
	mar := mar.NewMAR(dbus, rbus, cbus.NLm)
	ram := ram.NewRam(dbus, rbus, cbus.NCE)
	ir := ir.NewIR(dbus, ibus, cbus.NLi, cbus.NEi)
	cntl := cntl.NewCntl(clk, ibus, cbus)

	acc := acc.NewAcc(dbus, abus, cbus.NLa, cbus.Ea)
	alu := alu.NewALU(abus, bbus, dbus, cbus.Su, cbus.Eu)
	breg := reg.NewReg(dbus, bbus, cbus.NLb)
	oreg := reg.NewReg(dbus, obus, cbus.NLo)
	dis := dis.New(obus)

	clk.SetHigh(cntl.Clk)
	clk.SetLow(cntl.NClk)

	abus.SetTrigger([]func(){alu.Exec})
	bbus.SetTrigger([]func(){alu.Exec})
	obus.SetTrigger([]func(){dis.Exec})
	rbus.SetTrigger([]func(){ram.Exec})
	ibus.SetTrigger([]func(){cntl.Exec})

	cntl.SetTrigger("CLK", alu.Exec)

	cntl.SetTrigger("CLK", mar.Clk)
	cntl.SetTrigger("CLK", ram.Exec)
	cntl.SetTrigger("CLK", ir.Clk)
	cntl.SetTrigger("CLK", acc.Clk)

	cntl.SetTrigger("CLK", breg.Clk)
	cntl.SetTrigger("CLK", oreg.Clk)

	cntl.SetTrigger("NCLK", pc.NClk)
	cntl.SetTrigger("CLR", ir.Clr)
	cntl.SetTrigger("NCLR", pc.NClr)

	ram.BurnRam()

	dbg := dbg.NewDbg(
		acc, alu, dis, pc, mar, ram, ir,
		cntl, cbus, breg, oreg, abus, bbus, dbus, obus, rbus, ibus,
	)

	cntl.SetTrigger("CLK", dbg.Tick)
	cntl.SetTrigger("NCLK", dbg.Tick)

	return &Sap1{
		Clk:  clk,
		DBus: dbus, CBus: cbus,
		ABus: abus, BBus: bbus, OBus: obus,
		Acc: acc, Alu: alu, BReg: breg, OReg: oreg, Dis: dis,
		RBus: rbus, IBus: ibus,
		Pc: pc, Mar: mar, Ram: ram, Ir: ir, Cntl: cntl,
	}
}
