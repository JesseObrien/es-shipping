package main

import ()

var ATSEA = Port{Name: "At Sea"}

type EventLog struct {
	Events []DomainEvent
}

func (el *EventLog) Add(e DomainEvent) {
	el.Events = append(el.Events, e)
}

type ShippingProcessor struct {
	Log EventLog
}

func (sp *ShippingProcessor) Process() {

}

type EventProcessor interface {
	Process()
}

type DomainEvent interface {
	Process()
}

func main() {

}
