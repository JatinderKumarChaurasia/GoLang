package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("Value of I :", i)
		i = i + 1
	}
	for j := 4; j <= 7; j++ {
		fmt.Println("Value of J:", j)
	}
	for {
		fmt.Println("Loop with no args") //for without a condition will loop repeatedly until you break out of
		//the loop or return from the enclosing function.
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("ODD NUMBERS ARE:", n)
	}

	if num := -4; num < 0 { //HERE WE NEED TO USE SEMI-COLON TO DECLARE AND IT IS available to all branches of if -following
		fmt.Println("Number is Negative")
	} else if num > 0 {
		fmt.Println("Number is Positive")
	} else {
		fmt.Println("Number is e= TO 0")
	}

}