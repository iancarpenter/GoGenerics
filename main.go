package main

import (
	"fmt"
	"reflect"
)

// Generic function that works with any type
func HelloWorld[T any](value T) {
	fmt.Printf("Hello, %v! (Type: %s)\n", value, reflect.TypeOf(value))
}

func main() {
	HelloWorld("World") // Using a string
	HelloWorld(123)     // Using an integer
	HelloWorld(3.14)    // Using a float
	HelloWorld(true)    // Using a boolean
}
