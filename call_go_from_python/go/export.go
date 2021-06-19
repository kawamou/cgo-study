package main

import (
	"C"
	"fmt"
)

//export my_print
func my_print() {
	fmt.Println("Hello, World!")
}

func main() {}
