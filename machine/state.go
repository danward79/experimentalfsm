package machine

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/danward79/fsm/intList"
)

//State represents state for a given set of conditions
type State map[string]string

//mapStates turns a list of expected states into a map
func mapStates(file, ignoreColumns string) State {
	m := make(map[string]string)

	ignoreList, err := intList.New(ignoreColumns)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("State map file: ", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		var stateName, stateCondition string
		for k, v := range record {
			if k == 0 {
				stateName = v
			} else {
				if !ignore(ignoreList, int64(k)) {
					stateCondition += v
				}

			}
		}

		m[stateCondition] = stateName
	}
	return m
}
