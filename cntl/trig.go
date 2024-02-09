package cntl

type TrigType string

const (
	CLK  TrigType = "CLK"
	NCLK TrigType = "NCLK"
	CLR  TrigType = "CLR"
	NCLR TrigType = "NCLR"
)

func (cntl *Cntl) SetTrigger(tType TrigType, newTrig func()) {
	trigMap := map[TrigType]*[]func(){
		CLK:  &cntl.onClk,
		NCLK: &cntl.onNClk,
		CLR:  &cntl.onClr,
		NCLR: &cntl.onNClr,
	}

	if _, ok := trigMap[tType]; ok {
		*trigMap[tType] = append(*trigMap[tType], newTrig)
	}
}

func (cntl *Cntl) Trig(triggers []func()) {
	for _, trigger := range triggers {
		if trigger != nil {
			trigger()
		}
	}
}
