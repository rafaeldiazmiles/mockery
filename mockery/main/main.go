package main

import (
	"fmt"
)

// In order to implment Mock and to understand why we need the interface in the same package as the concrete type that will use it and why "we receive interfaces but return objects"

// MessageService handles notifying clients they have been charged
type MessageService interface {
	SendChargeNotification(int) error
}

// SMSService is our implementation of MessageService
type SMSService struct{}

// MyService uses the MessageService to notify clients
type MyService struct {
	messageService MessageService
}

// SendChargeNotification notifies clients they have been charged via SMS
func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending production charge notification")
	return nil
}

// ChargeCustomer performs the charge to the customer. In a real system we would maybe mock this as well but here, I want to make some money every time I run my tests
func (a MyService) ChargeCustomer(value int) error {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging customer for the value of %d\n", value)
	return nil
}

func main() {
	fmt.Println("Hello World")

	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}
