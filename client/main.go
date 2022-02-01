package main

import (
	"fmt"

	"github.com/go-object-oriented/payment"
)

/* Best practice:
If you working with other 3rd party library even they defined interface within their library
still copy interface into your own consumer package (not receiver package)
this way prevent to import package only because of you need to use interface type inside your consumer logic
furthermore this also easier to remove the imported 3rd party library when you decide not to use them anymore
So always define interface inside consumer package: where you need to use them, not inside the receiver package: where implement detail of interface
*/
type PaymentOption interface {
	ProcessPayment(float32) bool
}

func main() {
	// Inheritance sample
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

	// Interface sample (encapsulation + polymorphism)
	fmt.Println("use interface inheritance")
	var option PaymentOption
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

	// Message passing sample
	chargeCh := make(chan float32)
	payment.CreateDebitAccount(chargeCh)
	chargeCh <- 500
	// prevernt early exit wait for console input defore end the app
	var a string
	fmt.Scanln(&a)
}
