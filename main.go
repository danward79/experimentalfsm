package main

import (
	"fmt"
	"log"

	"github.com/danward79/fsm/csvStreamer"
	"github.com/danward79/fsm/machine"
)

func main() {

	//Create machine and output state map
	camShaft := machine.New("", "states/states.csv", "", "1")
	fmt.Println(camShaft)
	//Sense check state map
	/*for k, v := range camShaft.Output.state {
		fmt.Printf("k: %s, v: %s\n", k, v)
	}*/

	tractionOutStream, err := csvStream.New("sampledata/sample.csv", "0,38-44,47-49,51-52,51,52,54-60,63,65-67,69,79,80-82")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tractionOutStream)

	go tractionOutStream.Emit()

	for msg := range tractionOutStream.Out {
		fmt.Println(msg)
	}

}
