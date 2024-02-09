package clk

import (
	"time"
)

type Clock struct {
	Interval    time.Duration
	Clk         func()
	NClk        func()
	TickCounter int
	MaxTicks    int
	ClockState  bool
}

// NewClock: Constructs a new instance of Clock
func NewClock(maxTicks int) *Clock {
	return &Clock{MaxTicks: maxTicks, ClockState: false}
}

func (clock *Clock) SetHigh(Clk func()) { clock.Clk = Clk }
func (clock *Clock) SetLow(NClk func()) { clock.NClk = NClk }
func (clock *Clock) StopClk()           { clock.MaxTicks = 0 }

// StartClk starts the clock ticking
func (clock *Clock) StartClk() {
	go func() {
		ticker := time.NewTicker(Duration)
		defer ticker.Stop()
		for range ticker.C {
			if clock.TickCounter >= clock.MaxTicks {
				break
			}
			clock.onTick()
		}
	}()
}

// onTick handles the clock ticks,
// toggling between Clk and NClk methods
func (clock *Clock) onTick() {
	if clock.ClockState {
		clock.Clk()
	} else {
		clock.NClk()
	}
	clock.ClockState = !clock.ClockState
	clock.TickCounter++
}
