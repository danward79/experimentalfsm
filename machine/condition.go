package machine

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/danward79/fsm/intList"
)

//Condition type, represents a set of conditions which return a state
type Condition struct {
	file         string
	ignoreFields string
	State        map[string]string //ie. map[condition]state
}

//String pretty output for Condition types
func (c *Condition) String() string {
	fmt.Println("ssdfsdfsdf")
	var s string
	for k, v := range c.State {
		s += fmt.Sprintf("Condition: %s, State: %s\n", k, v)
	}
	return fmt.Sprintf("File: %s\nIgnore Fields: %s\nState Conditions: %s", c.file, c.ignoreFields, s)
}

//GetState current state
func (c *Condition) GetState(conditions string) string {
	return c.State[conditions]
}

//MapConditions map condition set to state set
func mapConditions(file, ignoreFields string) map[string]string {
	m := make(map[string]string)

	ignoreList, err := intList.New(ignoreFields)
	if err != nil {
		log.Fatal("Condition map file: ", err)
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		var stateName, stateValue string
		for k, v := range record {
			if k == 0 {
				stateName = v
			} else {
				if !ignore(ignoreList, int64(k)) {
					stateValue += v
				}
			}
		}

		m[stateValue] = stateName
	}
	return m
}
