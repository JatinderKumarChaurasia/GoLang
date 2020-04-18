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

	p1 := organization.NewPerson("Shivani", "Sharma", organization.NewSocialSecurityNumber("566-245-1223"))

	//err := p.SetTwitterHandler("@twitter")
	err := p.SetTwitterHandler(organization.TwitterHandler("@twitter"))
	fmt.Printf("%T\n", organization.TwitterHandler("String"))
	if err != nil {
		fmt.Printf("Error Occurred during setting twitter handler : %s", err.Error())
	}
	err1 := p1.SetTwitterHandler(organization.TwitterHandler("@twitter"))
	fmt.Printf("%T\n", organization.TwitterHandler("String"))
	if err1 != nil {
		fmt.Printf("Error Occured during setting twitter handler : %s", err1.Error())
	}

	fmt.Println("P1 Twitter Handler : ", p1.TwitterHandler())
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println("Person : ", p.FullName())

	/*
		// Testing Equality
		//name1 := Name{FirstName: "Shivani", LastName: "Sharma"}
		//name2 := OtherName{FirstName: "Shivani", LastName: "Sharma"}
		// name1 not equal to name2

		ssn := organization.NewSocialSecurityNumber("543-595-2848")
		eui := organization.NewEuropeanUnionIdentifier("022","Italy")

		if ssn == eui {
			fmt.Printf("Both are equal")
		} else {
			fmt.Println("Both are not equal")
		}
		name1 := Name{FirstName: "", LastName: ""}
		if name1 == (Name{}) {
			fmt.Println("We match")
		}
		name2 := &Name{}
		name2 = nil
		if name2 == nil {
			fmt.Println("We match")
		}

		portfolio := map[Name][]organization.Person{}
		portfolio[name1] = []organization.Person{*p}

		if name1 == (Name{}){
			fmt.Println("We match")
		}

	*/
}

// Comparing structures
/*
//Name != OtherName
type Name struct {
	FirstName string
	LastName  string
}


type OtherName struct {
	FirstName string
	LastName  string
}

*/
