package greetings

// Declare a greetings package to related functions

import "fmt"

// Hello retrns agreetings for the named person.
func Hello(name string) string {
	// Hello: Function name
	// name: parameter name
	// string(inside): parameter type
	// string(outside): return type
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// ":=" can declare & initialize a variable in one line
	// The equivilant is as follows,
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return messge
}
