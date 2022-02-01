package main

import "fmt"

type Account struct{}

func (c *Account) AvailableFunds() float32 {
	fmt.Println("Listing available fund...")
	return 0
}

func (c *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment...")
	return true
}

// Create way Listing Type without add field name, outside don't know CreditAccount have a Account's behivour
type CreditAccount struct {
	Account
}

// Compare this version, outside package knowing CreditAccount have a account type field
// type CreditAccount struct {
// 	account Account
// }

func main() {
	// composition sample:
	ca := &CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)
}
