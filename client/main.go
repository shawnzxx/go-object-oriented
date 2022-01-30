package main

import (
	"fmt"

	"github.com/go-object-oriented/payment"
)

func main() {
	// use class inheritance
	fmt.Println("use class inheritance")
	credit := payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	fmt.Printf("Owner name: %v\n", credit.OwnerName())
	credit.SetOwnerName("Shawn Zhang")
	fmt.Printf("New Owner name: %v\n", credit.OwnerName())
	fmt.Printf("Card number: %v\n", credit.CardNumber())
	fmt.Println("Trying to change card number")
	err := credit.SetCardNumber("invalid")
	if err != nil {
		fmt.Printf("SetCardNumber error: %v\n", err)
	}
	fmt.Printf("Available credit: %v\n", credit.AvailableCredit())

	// use interface inheritance
	fmt.Println("use interface inheritance")
	var option payment.PaymentOption
	// since CreditCard class implement the PaymentOption interface so can use interface as CreateCredit
	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	option.ProcessPayment(500)
	// since Cash class implement the PaymentOption interface so can use interface as Cash
	option = payment.CreateCashAccount()

	option.ProcessPayment(500)
}
