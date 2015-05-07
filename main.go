package main

import (
	"fmt"

	pdf "./pdf"
)

func main() {
	p := pdf.PDF{"bar"}
	fmt.Printf("%s", p.A)

}
