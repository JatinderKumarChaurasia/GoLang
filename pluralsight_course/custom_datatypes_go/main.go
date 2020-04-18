package main

import (
	"fmt"
	"github.com/pluralsight-custom-datatypes/customdatatypes/organization"
)

func main() {
	//p := organization.Person{}
	//var p organization.Identifiable = organization.Person{FirstName: "Shivani",LastName: "Sharma"} // still we can't access FirstName and LastName
	//p := organization.Person{FirstName: "Shivani", LastName: "Sharma"}
	//fmt.Println(p.ID(), p.FirstName, p.LastName)

	// Using constructor
	//p := organization.NewPerson("Shivani", "Sharma",organization.NewSocialSecurityNumber("566-245-1223"))
	p := organization.NewPerson("Shivani", "Sharma", organization.NewEuropeanUnionIdentifier("566-245-1223", "Britain"))

	//err := p.SetTwitterHandler("@twitter")
	err := p.SetTwitterHandler(organization.TwitterHandler("@twitter"))
	fmt.Printf("%T\n", organization.TwitterHandler("String"))
	if err != nil {
		fmt.Printf("Error Occured during setting twitter handler : %s", err.Error())
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println("Person : ", p.FullName())
}
