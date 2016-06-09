package machine

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/danward79/fsm/intList"
)

//ignore helper func to test if column should be ignored
func ignore(list intList.List, column int64) bool {
	for _, v := range list {
		if v == column {
			return true
		}
	}

	return false
}

//recordToString helper func to convert []string to string
func recordToString(record []string, ignoreColumns string) string {
	ignoreList, err := intList.New(ignoreColumns)
	if err != nil {
		log.Fatal(err)
	}

	var c string
	for k, v := range record {
		if !ignore(ignoreList, int64(k)) {
			c += v
		}
	}

	return c
}

//mapStates turns a list of expected states into a map
func mapStates(file, ignoreColumns string) State {
	m := make(map[string]string)

	ignoreList, err := intList.New(ignoreColumns)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(file)
	if err != nil {
		//TODO: Sort out this error handling
		log.Fatal(err)
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
