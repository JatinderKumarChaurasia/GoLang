package payment

import "fmt"

type Account struct{}

func (account *Account) GetAvailableFunds() float64 {
	fmt.Println("Getting Available Funds ............................... ")
	return 0.0
}

func (account *Account) ProcessCardPayment(amount float64) bool {
	fmt.Println("Processing Account Payment ......... ............ ")
	return true
}

type CreditAccount struct {
	Account
}

func (credit *CreditAccount) GetAvailableFunds() float64 {
	fmt.Println("Getting Credit Account Available Funds ............................... ")
	return 0.0
}

type CheckingAccount struct{}

func (checking CheckingAccount) GetAvailableFunds() float64 {
	fmt.Println("Getting Checking Account Available Funds ............................... ")
	return 0.0
}

type HybridAccount struct {
	CreditAccount
	CheckingAccount
}
