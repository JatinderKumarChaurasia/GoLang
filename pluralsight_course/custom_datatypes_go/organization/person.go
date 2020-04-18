package organization

import (
	"errors"
	"fmt"
	"strings"
)

//interface
type Identifiable interface {
	ID() string
}

type Handler struct {
	handle string
	name   string
}

type Citizen interface {
	Identifiable
	Country() string
}

type Name struct {
	firstName string
	lastName  string
}

// Adding Social Security Numbers
type socialSecurityNumber string

// constructor for SocialSecurityNumber
func NewSocialSecurityNumber(ssn string) Citizen {
	return socialSecurityNumber(ssn)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "World"
}

// For European Union Identifier
type europeanUnionIdentifier struct {
	europeanUnionId string
	country         string
}

func NewEuropeanUnionIdentifier(europeanUnionId, country string) Citizen {
	return europeanUnionIdentifier{
		europeanUnionId: europeanUnionId,
		country:         country,
	}
}

func (eu europeanUnionIdentifier) ID() string {
	return eu.europeanUnionId
}

func (eu europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("Country : %s", eu.country)
}

// type TwitterHandler = string // type alias // copy the fields and method sets // not able to extend the type
type TwitterHandler string // type declaration => copy the fields and the objects

func (handler TwitterHandler) RedirectURL() string {
	clearPrefixHandler := strings.TrimPrefix(string(handler), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", clearPrefixHandler)
}

type Employee struct {
	Name
}

// Struct
type Person struct {
	// Visible to Everyone -> Public Mode
	//FirstName string
	//LastName  string

	// private mode
	//firstName      string
	//lastName       string
	//name Name
	Name
	twitterHandler TwitterHandler
	Citizen
}

// go does not support constructor
//so ...
func NewPerson(firstName, lastName string, citizen Citizen) *Person {
	return &Person{
		//name: Name{firstName: firstName,lastName: lastName},
		Name:    Name{firstName: firstName, lastName: lastName},
		Citizen: citizen,
	}
}

//func (person *Person) FullName() string {
func (name *Name) FullName() string {
	//return fmt.Sprintf("%s %s", person.name.firstName, person.name.lastName)
	return fmt.Sprintf("%s %s", name.firstName, name.lastName)

}

func (person *Person) ID() string {
	//return "123456"
	return fmt.Sprintf("Social Security Number : %s", person.Citizen.ID())
}

func (person *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		person.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with @ symbol ")
	}
	person.twitterHandler = handler
	return nil
}

func (person *Person) TwitterHandler() TwitterHandler {
	return person.twitterHandler
}
