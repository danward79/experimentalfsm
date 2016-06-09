package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/danward79/fsm/csvStreamer"
	"github.com/danward79/fsm/machine"
)

func main() {
	inputFile := flag.String("d", "", "input data in csv form")
	columns := flag.String("dc", "", "columns to use")
	outputStates := flag.String("o", "", "output state data in csv form")
	outputIgnoreFields := flag.String("oc", "", "output state columns to ignore")
	inputStates := flag.String("i", "", "input state data in csv form")
	inputIgnoreFields := flag.String("ic", "", "input state columns to ignore")
	flag.Parse()

	//Create machine and output state map
	camShaft := machine.New(*inputStates, *outputStates, *inputIgnoreFields, *outputIgnoreFields)
	//sense check states
	fmt.Println(camShaft.Output)

	tractionOutStream, err := csvStream.New(*inputFile, *columns) //"0,38-44,47-49,51-52,51,52,54-60,63,65-67,69,79,80-82"
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(tractionOutStream)

	go tractionOutStream.Emit()

	//for msg := range tractionOutStream.Out {
	//	fmt.Println(msg)
	//}

}
