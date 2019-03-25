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

	panic("O my Good. Panic")
}
