package main

import "fmt"

func main() {

	a := Dog{Color: "Brown"}

	fmt.Println(a)

	// a := 7

	// if a > 10 {
	// 	array := make([]int, 5)
	// 	array[2] = 3
	// } else {
	// 	array := make([]string, 10)
	// 	array[5] = "sfsdf"
	// }

	// fmt.Println(array)

}

type Dog struct {
	Name     string
	YearsOld int
	Color    string
	Weight   float64
}

type SvSp struct {
	Left  int
	Right int
	Meta  string
}
