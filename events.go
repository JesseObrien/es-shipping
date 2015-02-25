package main

import (
	"time"
)

type EventTime struct {
	Occurred time.Time
	Recorded time.Time
}

type DepartureEvent struct {
	Time EventTime
	Port Port
	Ship Ship
}

func NewDepartureEvent(t time.Time, p Port, s Ship) *DepartureEvent {
	return &DepartureEvent{
		Time: EventTime{Occurred: t},
		Port: p,
		Ship: s,
	}
}

func (de *DepartureEvent) Process() {
	de.Ship.handleDeparture(de)
}

type ArrivalEvent struct {
	Time EventTime
	Port Port
	Ship Ship
}

func NewArrivalEvent(t time.Time, p Port, s Ship) *ArrivalEvent {
	return &ArrivalEvent{
		Time: EventTime{Occurred: t},
		Port: p,
		Ship: s,
	}
}

func (ae *ArrivalEvent) Process() {
	ae.Ship.handleArrival(ae)
}

type CargoEvent struct {
	Id   int
	Ship Ship
}

type LoadCargoEvent struct {
	Time  EventTime
	Cargo CargoEvent
}

func NewLoadCargoEvent(id int, t time.Time, s Ship) *LoadCargoEvent {
	return &LoadCargoEvent{
		Time:  EventTime{Occurred: t},
		Cargo: CargoEvent{Id: id, Ship: s},
	}
}

type UnloadCargoEvent struct {
	Time  EventTime
	Cargo CargoEvent
}

func NewUnloadCargoEvent(id int, t time.Time, s Ship) *UnloadCargoEvent {
	return &UnloadCargoEvent{
		Time:  EventTime{Occurred: t},
		Cargo: CargoEvent{Id: id, Ship: s},
	}
}
