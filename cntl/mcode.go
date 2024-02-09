package cntl

const (
	LDA = 0
	ADD = 1
	SUB = 2
	LDI = 3
	OUT = 14
	HLT = 15
)

const NOP uint16 = 0b001111100011

var Fetch = []uint16{
	0b010111100011, 0b101111100011, 0b001001100011,
}

var microPrograms = map[int][]uint16{
	LDA: {0b000110100011, 0b001011000011, NOP},
	ADD: {0b000110100011, 0b001011100001, 0b001111000111},
	LDI: {0b001110000011, NOP, NOP},
	SUB: {0b000110100011, 0b001011100001, 0b001111001111},
	OUT: {0b001111110010, NOP, NOP},
	HLT: {NOP, NOP, NOP},
}

// LoadMicroProg loads the relevant micro program.
func (cntl *Cntl) LoadMicroProg() {
	cmd := cntl.Ibus.Data

	if cmd == HLT {
		cntl.Clock.StopClk()
	}

	cmdProg := microPrograms[int(cmd)]
	cntl.MicroProg = append(Fetch, cmdProg...)
}

// NextMicro advances the micro program.
func (cntl *Cntl) NextMicro() {
	if cntl.MicroStep == 0 {
		cntl.MicroProg = Fetch
	}

	cntl.Data = cntl.MicroProg[cntl.MicroStep]
	cntl.Cbus.Write(cntl.Data)

	cntl.MicroStep = (cntl.MicroStep + 1) % 6
}
