package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type CardPaymentOption interface {
	ProcessCardPayment(float64) bool
}

type CreditCard struct {
	ownerName              string
	cardNumber             string
	expirationMonth        int
	expirationYear         int
	securityCode           int
	availableCreditBalance float64
}

// Card Number Pattern
var creditCardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

func NewCreditCardAccountOpen(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int ,chargeCH chan float64) *CreditCard {
	creditAccount := &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
	go func(chargeCH chan float64) {
		fmt.Println("Processing Card Payment ...... ... ")
		for amount := range chargeCH {
			fmt.Printf("For Amount : %v Procssing Payment ",amount)
			creditAccount.ProcessCardPayment(amount)
		}
	}(chargeCH)
	return creditAccount
}

func (creditCard CreditCard) GetCreditCardOwnerName() string {
	return creditCard.ownerName
}

func (creditCard *CreditCard) SetCreditCardOwnerName(ownerName string) error {
	if len(ownerName) == 0 {
		return errors.New("invalid owner name provided")
	}
	creditCard.ownerName = ownerName
	return nil
}

func (creditCard CreditCard) GetCreditCardNumber() string {
	return creditCard.cardNumber
}

func (creditCard *CreditCard) SetCreditCardNumber(cardNumber string) error {
	if !creditCardNumberPattern.Match([]byte(cardNumber)) {
		return errors.New("invalid credit card number format . pattern unmatched")
	}
	creditCard.cardNumber = cardNumber
	return nil
}

func (creditCard CreditCard) GetCreditCardExpirationDate() (int, int) {
	return creditCard.expirationMonth, creditCard.expirationYear
}

func (creditCard *CreditCard) SetCreditCardExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() || year == now.Year() && time.Month(month) < now.Month() {
		return errors.New("expiration date can not be smaller or equal to current")
	}
	creditCard.expirationMonth, creditCard.expirationYear = month, year
	return nil
}

func (creditCard CreditCard) GetCreditCardSecurityCode() int { return creditCard.securityCode }

func (creditCard *CreditCard) SetCreditCardSecurityCode(securityCode int) error {
	if securityCode < 100 && securityCode > 999 {
		return errors.New("invalid security code")
	}
	creditCard.securityCode = securityCode
	return nil
}

func (creditCard CreditCard) GetAvailableCreditCardBalance() float64 {
	return 50000.0
}

// Payment Process Methods
func (creditCard *CreditCard) ProcessCardPayment(amount float64) bool  {
	fmt.Println("Processing Credit Card Payment ............")
	return true
}