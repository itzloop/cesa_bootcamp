package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Constants for device states and events
const (
	// Light States
	StateOff    State = "Off"
	StateOn     State = "On"
	StateDimmed State = "Dimmed"

	// Thermostat States
	StateThermoOff State = "ThermoOff"
	StateHeating   State = "Heating"
	StateCooling   State = "Cooling"

	// Security System States
	StateDisarmed  State = "Disarmed"
	StateArmedStay State = "ArmedStay"
	StateArmedAway State = "ArmedAway"
	StateTriggered State = "Triggered"

	// Light Events
	EventTurnOn  Event = "TurnOn"
	EventTurnOff Event = "TurnOff"
	EventDim     Event = "Dim"

	// Thermostat Events
	EventStartHeating  Event = "StartHeating"
	EventStartCooling  Event = "StartCooling"
	EventTurnThermoOff Event = "TurnThermoOff"

	// Security System Events
	EventArmStay Event = "ArmStay"
	EventArmAway Event = "ArmAway"
	EventDisarm  Event = "Disarm"
	EventTrigger Event = "Trigger"
)

// Setup FSMs for devices
func setupLightFSM() FSM {
	// TODO: Implement based on the scenario which is given to you in markdown file
	return NewFSM("", nil)
}

func setupThermostatFSM() FSM {
	return NewFSM("", nil)
}

func setupSecuritySystemFSM() FSM {
	return NewFSM("", nil)
}

func TestLightFSM_BasicTransitions(t *testing.T) {
	lightFSM := setupLightFSM()

	err := lightFSM.TriggerEvent(EventTurnOn)
	assert.NoError(t, err)
	assert.Equal(t, StateOn, lightFSM.CurrentState())

	err = lightFSM.TriggerEvent(EventDim)
	assert.NoError(t, err)
	assert.Equal(t, StateDimmed, lightFSM.CurrentState())

	err = lightFSM.TriggerEvent(EventTurnOff)
	assert.NoError(t, err)
	assert.Equal(t, StateOff, lightFSM.CurrentState())
}

func TestThermostatFSM_BasicTransitions(t *testing.T) {
	thermoFSM := setupThermostatFSM()

	err := thermoFSM.TriggerEvent(EventStartHeating)
	assert.NoError(t, err)
	assert.Equal(t, StateHeating, thermoFSM.CurrentState())

	err = thermoFSM.TriggerEvent(EventTurnThermoOff)
	assert.NoError(t, err)
	assert.Equal(t, StateThermoOff, thermoFSM.CurrentState())

	err = thermoFSM.TriggerEvent(EventStartCooling)
	assert.NoError(t, err)
	assert.Equal(t, StateCooling, thermoFSM.CurrentState())
}

func TestSecuritySystemFSM_BasicTransitions(t *testing.T) {
	securityFSM := setupSecuritySystemFSM()

	err := securityFSM.TriggerEvent(EventArmStay)
	assert.NoError(t, err)
	assert.Equal(t, StateArmedStay, securityFSM.CurrentState())

	err = securityFSM.TriggerEvent(EventTrigger)
	assert.NoError(t, err)
	assert.Equal(t, StateTriggered, securityFSM.CurrentState())

	err = securityFSM.TriggerEvent(EventDisarm)
	assert.NoError(t, err)
	assert.Equal(t, StateDisarmed, securityFSM.CurrentState())
}


func TestSmartHomeFSM_NightModeScenario(t *testing.T) {
	lightFSM := setupLightFSM()
	thermoFSM := setupThermostatFSM()
	securityFSM := setupSecuritySystemFSM()

	// Simulate entering Night Mode
	lightFSM.TriggerEvent(EventTurnOn)
	lightFSM.TriggerEvent(EventDim)

	thermoFSM.TriggerEvent(EventStartHeating)

	securityFSM.TriggerEvent(EventArmStay)

	// Validate states
	assert.Equal(t, StateDimmed, lightFSM.CurrentState())
	assert.Equal(t, StateHeating, thermoFSM.CurrentState())
	assert.Equal(t, StateArmedStay, securityFSM.CurrentState())

	// Trigger intrusion
	err := securityFSM.TriggerEvent(EventTrigger)
	assert.NoError(t, err)

	// Validate triggered state
	assert.Equal(t, StateTriggered, securityFSM.CurrentState())
	// Lights should stay dimmed even if security is triggered
	assert.Equal(t, StateDimmed, lightFSM.CurrentState())
}

func TestSmartHomeFSM_AwayModeScenario(t *testing.T) {
	lightFSM := setupLightFSM()
	thermoFSM := setupThermostatFSM()
	securityFSM := setupSecuritySystemFSM()

	// Simulate entering Away Mode
	lightFSM.TriggerEvent(EventTurnOff)
	thermoFSM.TriggerEvent(EventTurnThermoOff)
	securityFSM.TriggerEvent(EventArmAway)

	// Validate states
	assert.Equal(t, StateOff, lightFSM.CurrentState())
	assert.Equal(t, StateThermoOff, thermoFSM.CurrentState())
	assert.Equal(t, StateArmedAway, securityFSM.CurrentState())
}

func TestSmartHomeFSM_ComplexScenario(t *testing.T) {
	lightFSM := setupLightFSM()
	thermoFSM := setupThermostatFSM()
	securityFSM := setupSecuritySystemFSM()

	// Start with Normal Mode
	lightFSM.TriggerEvent(EventTurnOn)
	thermoFSM.TriggerEvent(EventStartCooling)
	securityFSM.TriggerEvent(EventArmStay)

	// Validate initial states
	assert.Equal(t, StateOn, lightFSM.CurrentState())
	assert.Equal(t, StateCooling, thermoFSM.CurrentState())
	assert.Equal(t, StateArmedStay, securityFSM.CurrentState())

	// Simulate intrusion in normal mode
	err := securityFSM.TriggerEvent(EventTrigger)
	assert.NoError(t, err)

	// Validate system responses
	assert.Equal(t, StateTriggered, securityFSM.CurrentState())
	assert.Equal(t, StateOn, lightFSM.CurrentState()) // Lights remain on
}

func TestSmartHomeFSM_HeatInsteadOfCoolInNightMode(t *testing.T) {
	lightFSM := setupLightFSM()
	thermoFSM := setupThermostatFSM()
	securityFSM := setupSecuritySystemFSM()

	// Enter Night Mode
	lightFSM.TriggerEvent(EventTurnOn)
	lightFSM.TriggerEvent(EventDim)
	thermoFSM.TriggerEvent(EventStartHeating)
	securityFSM.TriggerEvent(EventArmStay)

	// Validate initial states
	assert.Equal(t, StateDimmed, lightFSM.CurrentState())
	assert.Equal(t, StateHeating, thermoFSM.CurrentState())
	assert.Equal(t, StateArmedStay, securityFSM.CurrentState())

	// Attempt to start cooling in night mode (invalid scenario)
	err := thermoFSM.TriggerEvent(EventStartCooling)
	assert.Error(t, err, "Cooling should not be allowed in Night Mode")
}


func TestLightFSM_WrongTransition(t *testing.T) {
	lightFSM := setupLightFSM()

	// Wrong transition: Trying to dim when off (shouldn't be allowed)
	err := lightFSM.TriggerEvent(EventDim)
	assert.NoError(t, err)  // Fix: This should expect an error
	assert.Equal(t, StateOff, lightFSM.CurrentState())
}

func TestThermostatFSM_StartHeatingWithoutTurningOff(t *testing.T) {
	thermoFSM := setupThermostatFSM()

	// Invalid: Start heating when already heating (shouldn't change state)
	thermoFSM.TriggerEvent(EventStartHeating)
	err := thermoFSM.TriggerEvent(EventStartHeating)
	assert.NoError(t, err)  // Fix:
}