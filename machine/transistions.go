package machine

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

//Transistions represents allowed transistions
type Transistions map[string]string

//Allowed is the transistion allowed
func (t Transistions) Allowed(from, to string) bool {
	return t[from] == to
}

//mapTransistion map transistion set to allowed state set
func mapTransistion(file string) map[string]string {
	m := make(map[string]string)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Transistion map file: ", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		m[record[0]] = record[1]
	}
	return m
}
