package main

/*
   if variable is new than use :=
   otherwise to update the previous declared variable use =
   function can return mutliple values
   a numeric constant has no type until it is given through explicit conversion
*/
import (
	"fmt"
	"math"
	"os"
)

const checkconst = "Learning Constant"

func main() {
	var name string
	//var age int
	//age =45
	//other way age:=getAge()
	//shorter way
	var a string = "Mutkl"
	fmt.Println(a)
	fmt.Println("Printing constant string:", checkconst)
	const intconst = 5 //A numeric constant has no type
	converconstint := int64(intconst)
	fmt.Println("Converted numeric constant to int64:", converconstint)
	fmt.Println("Using Math.sin():", math.Sin(intconst)) //using math
	var b, c int = 4, 5
	fmt.Println(b, c)

	education := "BSLC"
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Time is Over", os.Args[1])
	fmt.Println("Hello, playground", education, name, getAge())
	sof, ware := getMessage("boolean")
	fmt.Println("vALUE OF sOFT", sof, ware)
	_, was := getMessage("Only get one")
	fmt.Println("vALUE OF sOFT", was)
	pre, _ := getMessage("Only get one")
	fmt.Println("vALUE OF sOFT", pre)

}
func getAge() int {
	return 45
}
func getMessage(arfu string) (string, bool) {
	if len(arfu) >= 10 {
		return arfu, false
	} else {
		return arfu, true
	}
}