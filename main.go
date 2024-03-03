package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolen
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses base on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n",
		"tea", 1, 'A', 3.14, true, 2, 5)
}
