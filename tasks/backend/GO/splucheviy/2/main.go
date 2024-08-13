package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Chumba, dlya S vvedi 2 storony formata float")
		return
	}

	a, first_err := strconv.ParseFloat(args[1], 64)
	b, second_err := strconv.ParseFloat(args[2], 64)

	if first_err != nil || second_err != nil {
		fmt.Println(first_err)
		fmt.Println("Chumba, vvedi float")
		return
	}

	s := a * b
	fmt.Printf("S pryamougolnika: %f\n", s)
}
