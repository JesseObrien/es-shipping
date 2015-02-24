package main

type Ship struct {
	Name  string
	Port  Port
	Cargo map[string]Cargo
}

func (s *Ship) handleDeparture(de *DepartureEvent) {
	s.Port = ATSEA
}

func (s *Ship) handleArrival(ev *ArrivalEvent) {

	s.Port = ev.Port

	for _, c := range s.Cargo {
		c.handleArrival(ev)
	}

}
