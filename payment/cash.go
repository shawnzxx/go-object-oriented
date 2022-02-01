package payment

import "fmt"

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}

// Message passing through ProcessPayment interface
// outside can invoke this method
func (c *Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Procssing a cash transaction...")
	return true
}
