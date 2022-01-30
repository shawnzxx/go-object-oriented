package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

// Encapsulation1: external package can not access private fileds
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// Encapsulation2: external package use constructor to create CreditCard
func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// CreditCard class implement the interface
func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment...")
	return true
}

// Encapsulation3: external package use getter and setter to interact with private fileds
func (c *CreditCard) OwnerName() string {
	return c.ownerName
}

func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owner name provided")
	}
	c.ownerName = value
	return nil
}

func (c *CreditCard) CardNumber() string {
	return c.cardNumber
}

// CreditCard number validation must use - to link 4 digits
var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

func (c *CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card number format")
	}
	c.cardNumber = value
	return nil
}

func (c *CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

func (c *CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() ||
		(year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must lie in the future")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

func (c *CreditCard) SecurityCode() int {
	return c.securityCode
}

func (c *CreditCard) SetSecurityCode(value int) error {
	if value < 100 || value > 999 {
		return errors.New("CCV must bigger than 100 smaller than 999")
	}
	return nil
}

func (c *CreditCard) AvailableCredit() float32 {
	return 5000. // this can come from a web service, client doesn't know or care
}
