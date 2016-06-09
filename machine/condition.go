package machine

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/danward79/fsm/intList"
)

//Condition type, represents a set of conditions which return a state
type Condition struct {
	file         string
	ignoreFields string
	state        map[string]string //map[condition(string)]state(string)
}

//GetState current state
func (c *Condition) GetState(conditions string) string {
	return c.state[conditions]
}

//MapConditions map condition set to state set
func mapConditions(file, ignoreFields string) map[string]string {
	m := make(map[string]string)

	ignoreList, err := intList.New(ignoreFields)
	if err != nil {
		log.Fatal(err)
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
