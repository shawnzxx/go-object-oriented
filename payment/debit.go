package payment

import "fmt"

type DebitAccount struct{}

// private method outside no way to interactive with it
func (d *DebitAccount) processPayment(amount float32) {
	fmt.Println("Processing debit card payment...")
}

// All interaction through channel, from constructor param
func CreateDebitAccount(chargeCh chan float32) *DebitAccount {
	debit := &DebitAccount{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			debit.processPayment(amount)
		}
	}(chargeCh)
	return debit
}
