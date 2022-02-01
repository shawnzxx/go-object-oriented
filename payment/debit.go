package payment

import "fmt"

type DebitAccount struct{}

// outside can not invoke this method
func (d *DebitAccount) processPayment(amount float32) {
	fmt.Println("Processing debit card payment...")
}

// Message passing through Channel
// All interaction through channel, from constructor param
func CreateDebitAccount(chargeCh chan float32) *DebitAccount {
	debit := &DebitAccount{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			// internally we call this private method
			debit.processPayment(amount)
		}
	}(chargeCh)
	return debit
}
