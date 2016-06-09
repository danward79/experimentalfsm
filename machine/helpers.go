package machine

import (
	"log"

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

	var s string
	for k, v := range record {
		if !ignore(ignoreList, int64(k)) {
			s += v
		}
	}
	return s
}
