package dis

import "sap-1/bus"

/*
   Dis Simulates an 8-bit linear display.

         Obus (8)
            ▼
    ┏━━━━━━━┷━━━━━━━━┓
    ┃    Display     ┃
    ┃   ▢▢▢▢ ▢▢▢▢    ┃
    ┗━━━━━━━━━━━━━━━━┛
*/
type Dis struct {
	Obus *bus.Bus // Data (In) Bus

	Data uint8
}

// NewDis: Creates a new a-synchronous Display
func NewDis(Obus *bus.Bus) *Dis {
	return &Dis{Obus: Obus}
}

// Exec: Loads the data from the Obus into the Display
func (dis *Dis) Exec() {
	dis.Data = dis.Obus.Data
}
