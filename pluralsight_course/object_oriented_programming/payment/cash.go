package payment

import "fmt"

type CashAccount struct {}

func CreateCashAccount() *CashAccount{
	return &CashAccount{}
}

func(cashAccount *CashAccount) ProcessCardPayment(amount float64) bool {
	fmt.Println("Processing Cash Card Payment .... ")
	return true
}