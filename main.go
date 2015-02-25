package main

import (
	"time"
)

var ATSEA = Port{Name: "At Sea"}

const timeFmt = "Jan 2, 2006 at 3:04pm (MST)"

type EventLog struct {
	Events []DomainEvent
}

func (el *EventLog) Add(e DomainEvent) {
	el.Events = append(el.Events, e)
}

type ShippingProcessor struct {
	Log EventLog
}

func (sp *ShippingProcessor) Process(de DomainEvent) {
	de.Process()
	sp.Log.Add(de)
}

type EventProcessor interface {
	Process()
}

type DomainEvent interface {
	Process()
}

func main() {

	sp := ShippingProcessor{Log: EventLog{}}

	southampton := Port{Name: "Southampton", Country: "England"}
	cherbourg := Port{Name: "Cherbourg", Country: "France"}
	queenstown := Port{Name: "Queenstown", Country: "Ireland"}

	titanic := Ship{Name: "Titanic", Port: southampton}

	sht, _ := time.Parse(timeFmt, "Apr 10, 1912 at 12:00pm (UTC)")
	sp.Process(NewDepartureEvent(sht, southampton, titanic))
	cat, _ := time.Parse(timeFmt, "Apr 10, 1912 at 6:30pm (CET)")
	sp.Process(NewArrivalEvent(cat, cherbourg, titanic))
	cdt, _ := time.Parse(timeFmt, "Apr 10, 1912 at 8:10pm (CET)")
	sp.Process(NewDepartureEvent(cdt, cherbourg, titanic))
	qat, _ := time.Parse(timeFmt, "Apr 11, 1912 at 11:30am (UTC)")
	sp.Process(NewArrivalEvent(qat, queenstown, titanic))
	qdt, _ := time.Parse(timeFmt, "Apr 11, 1912 at 1:30pm (UTC)")
	sp.Process(NewDepartureEvent(qdt, queenstown, titanic))

}
