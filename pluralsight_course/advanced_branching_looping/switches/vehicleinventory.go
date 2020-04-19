package main

import "fmt"

type vehicle interface {
	printVehicleDetails(string)
}

type car struct {
	model, make, color, typeVehicle string
	cc                              float64
	year, seating, doors            int
}

type truck struct {
	model, make, color, typeVehicle string
	cc                              float64
	year, seating, doors            int
	loadingCapacity                 float64
}

type bike struct {
	model, make, color string
	cc                 float64
	year               int
}

var inventory []vehicle

func init() {

	var timeMachine vehicle

	inventory = []vehicle{
		bike{"FTR 1200", "Indian", "Black", 1203.0, 2019},
		bike{"Iron 1200", "Harley", "Blue", 1200.0, 2018},
		car{"Sonata", "Hyundai", "White", "Sedan", 2400.0, 2017, 5, 4},
		car{"SantaFe", "Hyundai", "Red", "SUV", 2400.0, 2016, 7, 4},
		car{"Civic", "Honda", "White", "Hatchback", 2000.0, 2017, 5, 4},
		car{"A5", "Audi", "Red", "Coupe", 3000.0, 2019, 2, 2},
		car{"Mazda6", "Mazda", "White", "Sedan", 2500.0, 2018, 5, 4},
		car{"CRV", "Honda", "Black", "SUV", 1500.0, 2017, 5, 4},
		car{"Camry", "Toyota", "Silver", "Sedan", 3500.0, 2018, 5, 4},
		truck{"F-150", "Ford", "Gray", "Truck", 3600.0, 2014, 7, 4, 13200.0},
		truck{"RAM1500", "Dodge", "White", "Truck", 5700.0, 2017, 5, 2, 12750.0},
		truck{"Silverado", "Chevrolet", "Black", "Truck", 6000.0, 2016, 7, 4, 14500.0},
		timeMachine,
	}
}

func main() {
	carCount, bikeCount, truckCount := 0, 0, 0
	fmt.Printf("%-15v%-12v%-12v%-10v%-4v%12v%12v%8v%20v\n", "Type", "Make", "Model", "Color", "CC", "Year", "Seating", "Doors", "Loading Capacity")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	for _, ve := range inventory {
		switch v := ve.(type) {
		case car:
			v.carDetails()
			carCount++
		case bike:
			v.bikeDetails()
			bikeCount++
		case truck:
			v.truckDetails()
			truckCount++
		default:
			fmt.Println("Vehicle type not exists !")
		}
	}
}

func (c *car) carDetails() {
	msg := fmt.Sprintf("Car-%-11v%-12v%-12v%-10v%-6.2f%9v%9d%9d\t        --\n", c.typeVehicle, c.make, c.model, c.color, c.cc, c.year, c.seating, c.doors)
	c.printVehicleDetails(msg)
}

func (c car) printVehicleDetails(msg string) {
	fmt.Printf("%v", msg)
}

func (b *bike) bikeDetails() {
	msg := fmt.Sprintf("Bike           %-12v%-12v%-10v%-6.2f%9v\t--\t --\t        --\n", b.make, b.model, b.color, b.cc, b.year)
	b.printVehicleDetails(msg)
}

func (b bike) printVehicleDetails(msg string) {
	fmt.Printf("%v", msg)
}

func (t *truck) truckDetails() {
	msg := fmt.Sprintf("Truck-%-9v%-12v%-12v%-10v%-6.2f%9v%9d%9d%20.2f\n", t.typeVehicle, t.make, t.model, t.color, t.cc, t.year, t.seating, t.doors, t.loadingCapacity)
	t.printVehicleDetails(msg)
}

func (t truck) printVehicleDetails(msg string) {
	fmt.Printf("%v", msg)
}
