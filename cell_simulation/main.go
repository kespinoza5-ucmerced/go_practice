package main

import (
	"fmt"
	"math/rand"

	"gitlab.com/akita/akita/v3/sim"
)

type CellSplitEvent struct {
	time    sim.VTimeInSec
	handler sim.Handler
}

func (e CellSplitEvent) Time() sim.VTimeInSec {
	return e.time
}

// Returns the handler that can should handle the event
func (e CellSplitEvent) Handler() sim.Handler {
	return e.handler
}

// IsSecondary tells if the event is a secondary event. Secondary event are
// handled after all same-time primary events are handled.
func (e CellSplitEvent) IsSecondary() bool {
	return false
}

type CellSplitEventHandler struct {
	cellCount int
	engine    sim.EventScheduler
}

func (h *CellSplitEventHandler) Handle(e sim.Event) error {
	// check type of event
	cellSplitEvent, ok := e.(CellSplitEvent)
	if !ok {
		panic("wrong event type")
	}

	// perform split
	h.cellCount++
	fmt.Printf("Cell split at time %f, cell count: %d\n",
		cellSplitEvent.time, h.cellCount)

	// randomly generate number between 1-2
	timeUntilNextSplit := sim.VTimeInSec(rand.Float64() + 1)

	// create an event that adds current time and timeUntilNextSplit
	event := CellSplitEvent{
		time:    cellSplitEvent.time + timeUntilNextSplit,
		handler: h,
	}

	// stop simulation at 200s
	if cellSplitEvent.time < 200 {
		h.engine.Schedule(event)
	}

	return nil
}

func main() {
	engine := sim.NewSerialEngine()
	handler := &CellSplitEventHandler{
		engine: engine,
	}
	event := CellSplitEvent{
		time:    0,
		handler: handler,
	}
	engine.Schedule(event)

	engine.Run()
}
