package main

import (
	"gitlab.com/akita/akita/v3/sim"
)

type CellSplitEvent struct {
	time    sim.VTimeInSec
	handler sim.Handler
}

type CellSplitEventHandler struct {
	something int
	engine    sim.Engine
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
