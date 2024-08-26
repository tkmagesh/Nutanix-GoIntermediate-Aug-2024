package main

import "fmt"

// Ver 1.0
type MessageGenerator struct {
}

func (mg *MessageGenerator) GetMessage() string {
	return "Have a nice day!"
}

func NewMessageGenerator() *MessageGenerator {
	return &MessageGenerator{}
}

/* Contract */
type IMessageGenerator interface {
	GetMessage() string
}

type Greeter struct {
	messageGenerator IMessageGenerator
}

func (g *Greeter) Greet(userName string) string {
	return fmt.Sprintf("Hi %s, %s", userName, g.messageGenerator.GetMessage())
}

func NewGreeter(mg IMessageGenerator) *Greeter {
	greeter := &Greeter{
		messageGenerator: mg,
	}
	return greeter
}

// ver 2.0
type GoodDayMessageGenerator struct {
}

func (gdmg *GoodDayMessageGenerator) GetMessage() string {
	return "Have a good day!"
}

func NewGoodDayMessageGenerator() *GoodDayMessageGenerator {
	return &GoodDayMessageGenerator{}
}

func main() {
	// tight coupling
	/*
		greeter := &Greeter{}
		fmt.Println(greeter.Greet("Magesh"))
	*/

	// loose coupling using contracts
	/*
		mg := NewMessageGenerator()
		greeter := NewGreeter(mg)
	*/

	// ver 2.0
	// use a different implementation for generating greet messages (GoodDayMessageGenerator)
	gdmg := NewGoodDayMessageGenerator()
	greeter := NewGreeter(gdmg)
	fmt.Println(greeter.Greet("Magesh"))
}
