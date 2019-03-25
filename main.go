package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if message := recover(); message != nil {
			fmt.Printf("Error message: '%v'\nStack trace :\n%v", message, string(debug.Stack()))
		}
	}()


	fmt.Println(division(12, 3))
	fmt.Println(division(12, 0))

	panic("O my Good. Panic")
}

func division (base, part int) int {
	if part == 0 {
		panic("Division by zero")
	}

	return  base / part
}
