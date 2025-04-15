package main

import (
	"fmt"
	"reflect"
)

// Generic function that swaps two values of the same type
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// Generic function that returns the type of the input value
func GetType[T any](value T) string {
	return reflect.TypeOf(value).String()
}

// Generic function that works with any type
func HelloWorld[T any](value T) {
	fmt.Printf("Hello, %v! (Type: %s)\n", value, reflect.TypeOf(value))
}

func main() {
	HelloWorld("World") // Using a string
	HelloWorld(123)     // Using an integer
	HelloWorld(3.14)    // Using a float
	HelloWorld(true)    // Using a boolean

	theType := GetType("Hello")
	fmt.Printf("The type of 'Hello' is: %s\n", theType)
	theType = GetType(42)
	fmt.Printf("The type of 42 is: %s\n", theType)

	a, b := 1, 2
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a, b = Swap(a, b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)

	s1, s2 := "Hello", "World"
	fmt.Printf("Before swap: s1 = %s, s2 = %s\n", s1, s2)
	s1, s2 = Swap(s1, s2)
	fmt.Printf("After swap: s1 = %s, s2 = %s\n", s1, s2)
}
