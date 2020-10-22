package main

import (
	"fmt"
	"strings"
)

func main() {

	// var name string = "Rajiv" // assigning type
	// nameT := "Rajiv"          // Go guesses the type, only works in funcs
	fmt.Println(multiply(2, 3))
	totalLength, upperName := LenToUpper("rajiv")
	fmt.Println(totalLength)
	// totalLength, _ := LenToUpper("rajiv") // _ is used to ignore values, if I only wanted to print one of the two values I'd use _ <-
	fmt.Println(totalLength, upperName)
	repeatMe("Rajiv", "Daniel", "Eric", "Seba", "Blue", "Furiel", "Straxel", "Annie", "Zinek")

}

func multiply(a, b int) int {

	return a * b

}

func LenToUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
