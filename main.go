package main

import (
	"fmt"
)

func main() {

	fmt.Println(Multiply(2, 3))
	totalLength, upperName := LenToUpper("rajiv")
	fmt.Println(totalLength, upperName)
	RepeatMe("Rajiv", "Daniel", "Eric", "Seba", "Blue", "Furiel", "Straxel", "Annie", "Zinek")
	fmt.Println(NakedUpperLen("Rajiv")) // naked return
	fmt.Println(foreachadd(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(foradd(1, 2, 5, 6, 2, 3, 4, 6, 7, 8, 2, 3, 4, 5, 6, 7, 8, 2, 3))

}
