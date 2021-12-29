package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// smsServiceMock
type smsServiceMock struct {
	mock.Mock
}

// Our mocked smssService method
func (m *smsServiceMock) SendChargeNotification(value int) error {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)

	// This records that the method was called and passes in the value it was called with

	// It then returns whatever we tell it to return. In this case true to simulate an SMS Service Notification sent out
	return nil
}

// We need to satisfy our MessageService interface which sadly means we have to stub out every method defined in that interface
func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// TestChargeCustomer is where the magic happens. Here we create our SMSService mock
func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)

	// We then define what should be returned form SendChargeNotification wen we pass in the value 100 to it. In this case, we want to return true as it was successful in sendin notification
	smsService.On("SendChargeNotification", 100).Return(true)

	// Next we want ot define the service we wish to test
	myService := MyService{smsService}
	// and call said method
	myService.ChargeCustomer(100)

	// At the end, we verify thatour myService.ChargeCustomer method called our mocked SendChargeNotification method
	smsService.AssertExpectations(t)
}
