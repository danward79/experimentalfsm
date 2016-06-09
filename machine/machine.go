package machine

import (
	"fmt"
	"log"
)

//Machine type
type Machine struct {
	Output       Condition //expected state for an output condition
	Input        Condition //expected state for an input condition
	Transistions Transistions
	CurrentState string
	PrevState    string
	NextState    string
	OutputData   chan []string
	InputData    chan []string
}

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
		m.Input.State = mapConditions(m.Input.file, m.Input.ignoreFields)
	}

	if m.Output.file != "" {
		m.Output.State = mapConditions(m.Output.file, m.Output.ignoreFields)
	}

	//TODO: Transistions file needs to not be hard coded.
	m.Transistions = mapTransistion("statedata/transistions.csv")

	return &m
}

//Ingest consumes a stream of condition data
func (m *Machine) Ingest(c chan []string) {
	first := true

	for record := range c {
		s := m.Match(recordToString(record, ""), m.Output.ignoreFields)

		if s != "" {

			if m.CurrentState == "" && first {
				log.Println("First Run setting prev, current, next, etc")
				m.CurrentState = s
				m.NextState = m.Transistions[s]
				first = false
				continue
			}

			//fmt.Println("P", m.PrevState, "C", m.CurrentState, "N", m.NextState)
			if m.NextState == s {
				//log.Println("Transistion allowed: ", m.CurrentState, "to", s)
				m.PrevState = m.CurrentState
				m.CurrentState = s
				m.NextState = m.Transistions[s]
			} else {
				log.Println("Transistion not allowed: ", m.CurrentState, "to", s)
				m.PrevState = m.CurrentState
				m.CurrentState = s
				m.NextState = m.Transistions[s]
			}
		}

	}
}

//Match an event to a states
func (m *Machine) Match(record, ignoreColumns string) string {
	if s, ok := m.Output.State[record]; ok {
		return s
	}
	log.Println("No state match for condition:", record)
	return ""
}
