package main

import "fmt"

// Observer interface
type Observer interface {
	Update(string)
}

// Concrete Observer
type EmailClient struct {
	name string
}

func (e *EmailClient) Update(message string) {
	fmt.Printf("%s received message: %s\n", e.name, message)
}

// Subject
type NotificationService struct {
	observers []Observer
}

func (s *NotificationService) AddObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *NotificationService) NotifyAll(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

func main() {
	service := &NotificationService{}

	client1 := &EmailClient{name: "Client1"}
	client2 := &EmailClient{name: "Client2"}

	service.AddObserver(client1)
	service.AddObserver(client2)

	service.NotifyAll("New Product Launched!")
}
