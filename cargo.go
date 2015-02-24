package main

type Cargo struct {
	Name        string
	Destination Port
}

func NewCargo(n string, p Port) *Cargo {
	return &Cargo{Name: n, Destination: p}
}

func (c *Cargo) handleArrival(ev *ArrivalEvent) {
	if c.Destination == ev.Port {
		// Unload the cargo from the ship

	}
}
