package main

import (
	"fmt"
)

func main() {
	hx := []string{"Sn.No.", "T-REX", "X COORD", "Y COORD", "HEADID", "HE 23 GOP"}
	hx = fixHeader(hx)
	fmt.Println(hx)
}
