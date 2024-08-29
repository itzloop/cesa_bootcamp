### **Scenario: Multi-Device Smart Home System**

You are tasked with developing and testing an FSM for a multi-device smart home system that controls various appliances (e.g., lights, thermostat, security system). The FSM should manage different states of each device and handle complex interactions between them.

#### **FSM Requirements**
1. **Devices**: Each device (light, thermostat, security system) has its own set of states.
   - **Light**: `Off`, `On`, `Dimmed`.
   - **Thermostat**: `Off`, `Heating`, `Cooling`.
   - **Security System**: `Disarmed`, `Armed-Stay`, `Armed-Away`, `Triggered`.

2. **Global States**: The entire system has global states that influence the behavior of all devices:
   - **Normal Mode**: All devices operate independently.
   - **Away Mode**: Lights are `Off`, thermostat is set to `Off`, and the security system is `Armed-Away`.
   - **Night Mode**: Lights are `Dimmed`, thermostat is set to `Heating`, and the security system is `Armed-Stay`.

3. **Transitions**:
   - Devices can transition between their individual states based on user input or system commands.
   - The system can switch between global modes, triggering corresponding device state changes.
   - The security system can transition to `Triggered` if an intrusion is detected in any mode.

4. **Conditions**:
   - The thermostat should not transition to `Cooling` in `Night Mode`.
   - The lights should automatically transition to `Dimmed` when entering `Night Mode`.
   - If the security system is `Triggered`, all devices should be set to an `Alert` state (e.g., lights turn `On`).

#### **Testing Tasks**
1. **Test State Transitions for Individual Devices**:
   - Ensure that each device correctly transitions between its states based on user commands and system modes.
   - Test edge cases, such as attempting to transition the thermostat to `Cooling` in `Night Mode`.

2. **Test Global Mode Transitions**:
   - Verify that switching to `Away Mode` correctly sets all devices to their `Away` states.
   - Ensure that transitioning to `Night Mode` dims the lights, heats the house, and arms the security system.

3. **Test Inter-Device Dependencies**:
   - Test the scenario where the security system is triggered while in `Away Mode`. Ensure that all devices transition to their `Alert` states.
   - Ensure that the lights remain `Off` when transitioning from `Away Mode` to `Normal Mode` if the lights were previously `Off`.

