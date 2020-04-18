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

// type TwitterHandler = string // type alias // copy the fields and method sets // not able to extend the type
type TwitterHandler string // type declaration => copy the fields and the objects

func (handler TwitterHandler) RedirectURL() string {
	clearPrefixHandler := strings.TrimPrefix(string(handler), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", clearPrefixHandler)
}

type Employee struct {
	firstName string
	lastName  string
}

// Struct
type Person struct {
	// Visible to Everyone -> Public Mode
	//FirstName string
	//LastName  string

	// private mode
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

// go does not support constructor
//so ...
func NewPerson(firstName, lastName string) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (person *Person) FullName() string {
	return fmt.Sprintf("%s %s", person.firstName, person.lastName)
}

func (person *Person) ID() string {
	return "123456"
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
