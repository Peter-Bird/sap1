// pc/pc.go

package pc

import "sap-1/bus"

/*
	PC represents a 4-bit Program Counter
	interfacing with a system bus.

              ┏━━━━━━━━━━━━┓
      Cp ─────┨            ┃
    NCLK ─────┨▷    PC     ┠──── Wbus (4) ▶
    NCLR ─────┨    ▢▢▢▢    ┃
      Ep ─────┨            ┃
              ┗━━━━━━━━━━━━┛
*/

const fourBitMask = 0x0F // Mask for ensuring 4-bit data

// It supports operations like clear (NClr), clock (NClk),
// count (Cp), and enable (Ep) for data transfer.
type PC struct {
	wbus *bus.Bus // System bus connection

	cp func() bool // Count pulse, triggers PC increment
	ep func() bool // Enable pulse, triggers data transfer to bus

	data uint8 // Current counter value, 4-bit
}

// NewPC creates a new instance of a Program Counter.
func New(wbus *bus.Bus, cp, ep func() bool) *PC {
	return &PC{
		wbus: wbus,
		cp:   cp,
		ep:   ep,
		data: 0,
	}
}

// NClr resets the Program Counter to zero.
func (pc *PC) NClr() {
	pc.data = 0
}

// NClk processes the clock's negative edge, handling data transfer
// and counter increment based on control signals.
func (pc *PC) NClk() {
	if pc.ep() {
		pc.wbus.Write(pc.data & fourBitMask)
	}
	if pc.cp() {
		pc.data = (pc.data + 1) & fourBitMask
	}
}

// Read returns the value of the Program Counter
func (pc *PC) Read() uint8 {
	return pc.data
}

// End of package <PC>
