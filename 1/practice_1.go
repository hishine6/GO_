package main

/**
	Declaere a main package. Package is a way to group functions,
	and it's made up of all the files in the same directory
**/

import (
	"fmt"

	"rsc.io/quote"
)

/**
	Import the (popular) fmt package, which contains functions for formatting text,
	incluing printing to the console. This package is one of the standard library
	packages you got when you installed GO.
**/

/**
	After adding the "rsc.io/quote" package, you should run
	$ go mod tidy
	Go will add the quote module as a requirement, as well as a go.sum file
	for use in authenticating the moudle. In other words, it located and
	downloaded the rsc.io/quote module that contains the package you imported.
**/

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

/**
	Implement a main function to print a message to the console. A main
	function executes by default when you run the main package.
**/
