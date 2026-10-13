package main

import (
	"errors"
	"fmt"
	"strings"
)

type Notification struct {
	Recipient string
	Message string
}

func (n Notification) Summary() string {
	return n.Recipient + " : " + n.Message;
}

func (n Notification) Validate() error {
	if strings.TrimSpace(n.Recipient) == "" {
		return errors.New("Recipient is required in type Notification")
	} else if strings.TrimSpace(n.Message) == "" {
		return errors.New("Message is required in type Notification")
	}
	return nil
}

func main() {
	notification1 := Notification{
		Recipient : "John",
		Message : "You have a buyer",
	}
	notification2 := Notification{
		Recipient : "John",
		Message : " ",
	}
	notification3 := Notification{
		Recipient : "",
		Message : "You have a seller",
	}
	if error := notification1.Validate(); error == nil {
		fmt.Println("Validation for notification successful " + notification1.Summary())
	} else {
		fmt.Println("Validation failed : ", error)
	}
	if error := notification2.Validate(); error == nil {
		fmt.Println("Validation for notification successful " + notification2.Summary())
	} else {
		fmt.Println("Validation failed : ", error)
	}
	if error := notification3.Validate(); error == nil {
		fmt.Println("Validation for notification successful " + notification3.Summary())
	} else {
		fmt.Println("Validation failed : ", error)
	}
}