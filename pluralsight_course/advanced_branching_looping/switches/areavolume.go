package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0
	for {
		fmt.Println("\n ******************************** Main Menu *********************************")
		fmt.Print(" \n Choose any one of these : Area(1) Volume(2), Exit(0) : ")
		if _, err := fmt.Scan(&x); err == nil {
			switch x {
			case 1:
				y := 0
				for {
					fmt.Println("\n ************************ SubMenu ********* Calculate Area *****************")
					fmt.Print("\nChoose any one of these => Calculate Area for : Rectangle(1), Circle(2) ,Main Menu(9), Exit(0) : ")
					if _, err = fmt.Scan(&y); err == nil {
						switch y {
						case 1:
							areaRectangle()
						case 2:
							areaCircle()
						case 9:
							goto exitArea
						case 0:
							goto outerloop
						default:
							fmt.Println("Please enter the correct value for Submenu.Enter Again ")
						}
					}
				}
			exitArea:
				break
			case 2:
				y := 0
				for {
					fmt.Println("\n ************************ SubMenu ********* Calculate Volume *****************")
					fmt.Print("\nChoose any one of these => Calculate Volume for : Cylinder(1), Cube(2) ,Main Menu(9), Exit(0) : ")
					if _, err = fmt.Scan(&y); err == nil {
						switch y {
						case 1:
							volumeCylinder()
						case 2:
							volumeCube()
						case 9:
							goto exitVolume
						case 0:
							goto outerloop
						default:
							fmt.Println("Please enter the correct value for Submenu.Enter Again ")
						}
					}
				}
			exitVolume:
				break
			case 0:
				goto outerloop
				//os.Exit(0)
			default:
				fmt.Println("Please enter the correct value for Main Menu.Enter Again ")
			}
		}
	}
outerloop:
}

func volumeCube() {
	var edge float64
	fmt.Print("Enter the edge of cube : ")
	if _, err := fmt.Scan(&edge); err == nil {
		fmt.Printf("Volume of Cube for edge : %5.2f is : %5.2f", edge, edge*edge*edge)
	}
}

func volumeCylinder() {
	var radius, height float64
	fmt.Println("\n Enter the radius and height of cylinder : ")
	if _, err := fmt.Scan(&radius, &height); err == nil {
		fmt.Printf("Volume of cylinder for radius: %5.2f and height: %5.2f is : %5.2f\n", radius, height, math.Pi*radius*radius*height)
	}
}

func areaCircle() {
	var radius float64
	fmt.Print("\nEnter the Radius of Circle : ")
	if _, err := fmt.Scan(&radius); err == nil {
		fmt.Printf("The Area of Circle for Radius: %5.2f is : %5.2f\n", radius, math.Pi*radius*radius)
	}
}

func areaRectangle() {
	var length, width float64
	fmt.Println("\n Enter the Length and Width of Rectangle")
	if _, err := fmt.Scan(&length, &width); err == nil {
		fmt.Printf("Area of Rectangle for Length: %5.2f and Width: %5.2f is : %5.2f\n", length, width, length*width)
	}
}
