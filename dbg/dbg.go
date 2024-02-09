package dbg

import (
	"sap-1/acc"
	"sap-1/alu"
	"sap-1/bus"
	"sap-1/cntl"
	"sap-1/cntlBus"
	"sap-1/dis"
	"sap-1/ir"
	"sap-1/mar"
	"sap-1/pc"
	"sap-1/ram"
	"sap-1/reg"
	"sap-1/scr"
)

type Dbg struct {
	Acc  *acc.Acc
	Alu  *alu.ALU
	Dis  *dis.Dis
	Pc   *pc.PC
	Mar  *mar.MAR
	Ram  *ram.Ram
	Ir   *ir.IR
	Cntl *cntl.Cntl
	CBus *cntlBus.CntlBus
	Scr  *scr.Screen

	BReg, OReg *reg.Reg

	ABus, BBus, DBus, OBus, RBus, IBus *bus.Bus
}

// NewDbg: Creates a new instance of a Debug Module.
func NewDbg(
	Acc *acc.Acc, Alu *alu.ALU, Dis *dis.Dis, Pc *pc.PC, Mar *mar.MAR,
	Ram *ram.Ram, Ir *ir.IR, Cntl *cntl.Cntl, CBus *cntlBus.CntlBus,
	BReg, OReg *reg.Reg, ABus, BBus, DBus, OBus, RBus, IBus *bus.Bus,
) *Dbg {

	Scr := scr.NewScreen()

	return &Dbg{
		Acc: Acc, Alu: Alu, BReg: BReg, OReg: OReg, Dis: Dis, Pc: Pc, Mar: Mar,
		Ram: Ram, Ir: Ir, Cntl: Cntl, ABus: ABus, BBus: BBus, DBus: DBus,
		OBus: OBus, CBus: CBus, RBus: RBus, IBus: IBus, Scr: Scr,
	}
}

var count uint16 = 1

func (dbg *Dbg) Tick() {

	buff := scr.Buff{
		Wbus: dbg.DBus.Data, Abus: dbg.ABus.Data, Bbus: dbg.BBus.Data,
		Obus: dbg.OBus.Data, Rbus: dbg.RBus.Data, Ibus: dbg.IBus.Data,
		Cbus: dbg.CBus.Data, PC: dbg.Pc.Data, Acc: dbg.Acc.Data,
		Alu: dbg.Alu.Data, BReg: dbg.BReg.Data, OReg: dbg.OReg.Data,
		Dis: dbg.Dis.Data, Mar: dbg.Mar.Data, Ram: dbg.Ram.Data,
		Ir: dbg.Ir.Data, Cntl: dbg.Cntl.Data, Mic: dbg.Cntl.MicroProg, Cnt: count,
	}

	dbg.Scr.Tick(buff)

	count++
}
