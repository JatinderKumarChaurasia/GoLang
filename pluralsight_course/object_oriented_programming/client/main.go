package main

import (
	"github.com/pluralsight-course/objectoriented/payment"
)

func main() {

	/*
		creditCard := payment.NewCreditCardAccountOpen("Shivani Sharma","5000-5000-5000-5000",5,2022,345)
		fmt.Printf("Owner Name : %s\n",creditCard.GetCreditCardOwnerName())
		fmt.Printf("Card Number : %s\n",creditCard.GetCreditCardNumber())
		month,year := creditCard.GetCreditCardExpirationDate()
		fmt.Printf("Expiration Date : %d %d\n",month,year)
		fmt.Printf("Security Code : %v\n",creditCard.GetCreditCardSecurityCode())
		// Changing Security code
		err := creditCard.SetCreditCardNumber("invalid")
		if err != nil {
			fmt.Printf("Invalid Card Number : %v\n",err)
		}
		fmt.Printf("Available Credit Balance : %v",creditCard.GetAvailableCreditCardBalance())

	*/

	var option payment.CardPaymentOption

	chargeCh := make(chan float64)

	option = payment.NewCreditCardAccountOpen("Shivani Sharma", "5000-5000-5000-5000", 5, 2022, 345, chargeCh)
	chargeCh <- 500 //Pushing into channel
	option.ProcessCardPayment(500)
	option = payment.CreateCashAccount()
	option.ProcessCardPayment(500)

	ca := &payment.CreditAccount{}
	ca.ProcessCardPayment(500)
}
