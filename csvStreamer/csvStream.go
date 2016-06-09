package csvStream

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/danward79/fsm/intList"
)

//CSVStream type
type CSVStream struct {
	file    string
	columns intList.List
}

//String prints details about a CSVStream
func (c *CSVStream) String() string {
	return fmt.Sprintf("File: %s\nColumns: %s, Length: %d", c.file, c.columns, len(c.columns))
}

//New create a new CSVStream
func New(f string, cl string) (*CSVStream, error) {

	il, err := intList.New(cl)
	if err != nil {
		return nil, err
	}

	c := CSVStream{
		file:    f,
		columns: il,
	}

	return &c, nil
}

//Emit fields into output channel
func (c *CSVStream) Emit() chan []string {
	out := make(chan []string)

	go func(chanOut chan []string) {
		fi, err := os.Open(c.file)
		if err != nil {
			//TODO: Sort out this error handling
			log.Fatal(err)
		}
		defer fi.Close()

		csvReader := csv.NewReader(fi)

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}

			var r []string
			if len(c.columns) != 0 {
				for _, v := range c.columns {
					r = append(r, record[v])
				}
				record = r
			}

			chanOut <- record
		}
	}(out)

	return out
}

/*
zr, err := zip.OpenReader(c.file)
if err != nil {
	log.Fatal(err)
}
defer zr.Close()

for _, f := range zr.File {
	fmt.Printf("Contents of %s:\n", f.Name)
	r, err := f.Open()
	if err != nil {
		log.Fatal(err)
	}
	c.reader = csv.NewReader(r)

	for {
		record, err := c.reader.Read()
		if err == io.EOF {
			break
		}

		var r []string
		if len(c.columns) != 0 {
			for _, v := range c.columns {
				r = append(r, record[v])
			}
			record = r
		}

		c.Out <- record
	}
	r.Close()

}

*/
