package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	hx := []string{"Sn.No.", "T-REX", "X COORD", "Y COORD", "HEADID", "HE 23 GOP"}
	hx = fixHeader(hx)

	f, err := os.Open("testdata/test.csv")
	if err != nil {
		log.Fatalln("csv file can't be opened")
	}
	r := csv.NewReader(f)

	c := 0
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if c == 0 {
			rec = fixHeader(rec)
		}
		fmt.Println(rec)
		c++
	}

}
