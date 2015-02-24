package main

import (
	"time"
)

type BaseEvent struct {
	Occurred time.Time
	Recorded time.Time
}

type DepartureEvent struct {
	Base BaseEvent
	Port Port
	Ship Ship
}

func NewDepartureEvent(t time.Time, p Port, s Ship) *DepartureEvent {
	return &DepartureEvent{Base: BaseEvent{Occurred: t}, Port: p, Ship: s}
}

func (de *DepartureEvent) Process() {
	de.Ship.handleDeparture(de)
}

type ArrivalEvent struct {
	Time time.Time
	Port Port
	Ship Ship
}

func NewArrivalEvent(t time.Time, p Port, s Ship) *ArrivalEvent {
	return &ArrivalEvent{Time: t, Port: p, Ship: s}
}

func (ae *ArrivalEvent) Process() {
	ae.Ship.handleArrival(ae)
}

type LoadEvent struct {
	Id   int
	Time time.Time
	Ship Ship
}

func NewLoadEvent(id int, t time.Time, s Ship) *LoadEvent {
	return &LoadEvent{Id: id, Time: t, Ship: s}
}

type UnloadEvent struct {
	Id   int
	Time time.Time
	Ship Ship
}

func NewUnloadEvent(id int, t time.Time, s Ship) *UnloadEvent {
	return &UnloadEvent{Id: id, Time: t, Ship: s}
}
