package machine

import "fmt"

//Machine type
type Machine struct {
	Output Condition //expected state for an output condition
	Input  Condition //expected state for an input condition

	OutputData chan []string
	InputData  chan []string
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

	return &m
}

//Ingest consumes a stream of condition data
func (m *Machine) Ingest(c chan []string) {
	for record := range c {
		fmt.Println("Ingest record:", record)
		fmt.Println("match:", m.Match(record, m.Output.ignoreFields))
	}
}

//Match an event to a states
func (m *Machine) Match(record []string, ignoreColumns string) string {
	return m.Output.State[recordToString(record, "")]
}
