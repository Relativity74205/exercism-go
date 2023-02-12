package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greater interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct {
}

func (greater Italian) LanguageName() string {
	return "Italian"
}

func (greater Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {
}

func (greater Portuguese) LanguageName() string {
	return "Portuguese"
}

func (greater Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func SayHello(name string, greater Greater) string {
	return fmt.Sprintf("I can speak %s: %s", greater.LanguageName(), greater.Greet(name))
}
