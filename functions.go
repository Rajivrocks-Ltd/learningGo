package main

import (
	"fmt"
	"strings"
)

func Multiply(a, b int) int {

	return a * b

}

func LenToUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func RepeatMe(words ...string) {
	fmt.Println(words)
}

// Naked return
func NakedUpperLen(name string) (length int, uppercase string) {
	defer fmt.Println("I'm Done") // exec code after the function is finished
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func foreachadd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

func foradd(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		// fmt.Println(numbers[i])
		total += numbers[i]
	}
	return total
}
