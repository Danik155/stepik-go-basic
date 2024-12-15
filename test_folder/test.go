package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("test.txt", os.O_RDWR, 06)
	fmt.Println(err)
	if err != nil {
		file, err = os.Create("test.txt")
		fmt.Println(err)
	}

	_, err = file.WriteString("Haloe564645\n")
	fmt.Println(err)

	err = file.Close()

	fmt.Println(err)

}
