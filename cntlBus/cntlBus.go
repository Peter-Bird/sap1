package cntlBus

/*
	CntlBus:
	0 0 0 0  Cp Ep NLm NCE  NLi NEi NLa Ea  Su Eu NLb NLo

*/
type CntlBus struct{ Data uint16 }

// NewBus: Creates a new instance of the Control Bus.
func NewCntlBus() *CntlBus { return &CntlBus{} }

// Write: Sets the control signals on the bus.
func (bus *CntlBus) Write(value uint16) { bus.Data = value }

// Functions that enable reading individual signals from the bus.
func (bus *CntlBus) Cp() bool  { return bus.Signal(0xB) }
func (bus *CntlBus) Ep() bool  { return bus.Signal(0xA) }
func (bus *CntlBus) NLm() bool { return bus.Signal(0x9) }
func (bus *CntlBus) NCE() bool { return bus.Signal(0x8) }
func (bus *CntlBus) NLi() bool { return bus.Signal(0x7) }
func (bus *CntlBus) NEi() bool { return bus.Signal(0x6) }
func (bus *CntlBus) NLa() bool { return bus.Signal(0x5) }
func (bus *CntlBus) Ea() bool  { return bus.Signal(0x4) }
func (bus *CntlBus) Su() bool  { return bus.Signal(0x3) }
func (bus *CntlBus) Eu() bool  { return bus.Signal(0x2) }
func (bus *CntlBus) NLb() bool { return bus.Signal(0x1) }
func (bus *CntlBus) NLo() bool { return bus.Signal(0x0) }

// Signal: Returns the specified bit in the bus.
func (bus *CntlBus) Signal(bitIndex uint16) bool {
	return (bus.Data & (1 << bitIndex)) != 0
}
