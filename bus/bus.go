package bus

type Bus struct {
	Trigger []func()
	Data    uint8
}

// NewBus: Constructs a new instance of a Bus.
func NewBus() *Bus {
	return &Bus{Trigger: nil}
}

// SetTrigger: Sets the func to be executed when the bus changes.
func (bus *Bus) SetTrigger(newTriggers []func()) {
	bus.Trigger = append(bus.Trigger, newTriggers...)
}

/*
	Write: Sets the value to the bus's Data.
	It then triggers the recieving end that Data has changed.
*/
func (bus *Bus) Write(value uint8) {
	bus.Data = value

	if bus.Trigger != nil {
		for _, trigger := range bus.Trigger {
			trigger()
		}
	}
}
