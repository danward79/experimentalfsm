package machine

import (
	"fmt"

	"github.com/danward79/fsm/intList"
)

//Machine type
type Machine struct {
	Output Condition //expected state for an output condition
	Input  Condition //expected state for an input condition

	OutputData chan []string
	InputData  chan []string
}

//State machine state can be entered from a previous state
//type State map[string]string

//String print info about the machine
func (m *Machine) String() string {
	return fmt.Sprintf("Input: %s\n Output: %s\n", m.Input, m.Output)
}

//New generate a new machine and its events
func New(fileInputStates, fileOutputStates, inputColumnsIgnore, outputColumnsIgnore string) *Machine {
	m := Machine{
		Input: Condition{
			file:         fileInputStates,
			ignoreFields: inputColumnsIgnore,
		},
		Output: Condition{
			file:         fileOutputStates,
			ignoreFields: outputColumnsIgnore,
		},
	}

	if m.Input.file != "" {
		m.Input.state = mapConditions(m.Input.file, m.Input.ignoreFields)
	}

	if m.Output.file != "" {
		m.Output.state = mapConditions(m.Output.file, m.Output.ignoreFields)
	}

	return &m
}

//ignore helper func to test if column should be ignored
func ignore(list intList.List, column int64) bool {
	for _, v := range list {
		if v == column {
			return true
		}
	}

	return false
}
