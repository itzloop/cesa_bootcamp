package main

import (
	"errors"
	"fmt"
)

// State represents the state of the FSM
type State string

// Event represents an event that triggers a state transition
type Event string

// FSM interface defines the basic methods for a finite state machine
type FSM interface {
	CurrentState() State
	TriggerEvent(event Event) error
	AvailableTransitions() []Event
}

// GeneralFSM is a general implementation of the FSM interface
type GeneralFSM struct {
	currentState State
	transitions  map[State]map[Event]State
}

// NewFSM creates a new FSM with the initial state and transitions
func NewFSM(initialState State, transitions map[State]map[Event]State) *GeneralFSM {
	return &GeneralFSM{
		currentState: initialState,
		transitions:  transitions,
	}
}

// CurrentState returns the current state of the FSM
func (fsm *GeneralFSM) CurrentState() State {
	return fsm.currentState
}

// TriggerEvent triggers an event and transitions the FSM to a new state if possible
func (fsm *GeneralFSM) TriggerEvent(event Event) error {
	// Implement me
	nextState, ok := fsm.transitions[fsm.currentState][event]
	if !ok {
		return errors.New(fmt.Sprintf("Invalid transition: %s from state %s", event, fsm.currentState))
	}
	fsm.currentState = nextState
	return nil
}

// AvailableTransitions returns a list of events that can trigger a state transition from the current state
func (fsm *GeneralFSM) AvailableTransitions() []Event {
	// Implement me
	events := make([]Event, 0, len(fsm.transitions[fsm.currentState]))
	for event := range fsm.transitions[fsm.currentState] {
		events = append(events, event)
	}
	return events
}
