package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("Not enough arguments. Must enter the lengths of all sides.")
		return
	}

	a, err1 := strconv.ParseFloat(args[1], 64)
	b, err2 := strconv.ParseFloat(args[2], 64)
	c, err3 := strconv.ParseFloat(args[3], 64)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		fmt.Println("Invalid input. All arguments must be valid floating-point numbers.")
		return
	}

	if a+b < c || b+c < a || a+c < b {
		fmt.Println("This triangle does not exist")
		return
	}
	perimeter := a + b + c
	fmt.Printf("Perimeter of the triangle: %.2f\n", perimeter)

	semiPerimeter := perimeter / 2.0
	area := math.Sqrt(semiPerimeter * (semiPerimeter - a) * (semiPerimeter - b) * (semiPerimeter - c))
	fmt.Printf("Area of the triangle: %.2f\n", area)

	R := (a * b * c) / (4 * area)
	r := area / semiPerimeter

	angleAB := math.Asin(c/(2*R)) * (180.0 / math.Pi)
	angleBC := math.Asin(a/(2*R)) * (180.0 / math.Pi)
	angleAC := math.Asin(b/(2*R)) * (180.0 / math.Pi)

	fmt.Printf("Angle between a and b: %.2f\n", angleAB)
	fmt.Printf("Angle between b and c: %.2f\n", angleBC)
	fmt.Printf("Angle between a and c: %.2f\n", angleAC)
	fmt.Printf("Circumcircle radius: %.2f\n", R)
	fmt.Printf("radius of inscribed circle: %.2f\n", r)
}
