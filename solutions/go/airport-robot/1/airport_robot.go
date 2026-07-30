package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(Name string) string
}

func SayHello(Name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(Name))
}

type Italian struct {
	Name string
}

type Portuguese struct {
	Name string
}

func (it Italian) LanguageName() string {
	return "Italian"
}

func (pr Portuguese) LanguageName() string {
	return "Portuguese"
}

func (it Italian) Greet(Name string) string {
	return fmt.Sprintf("Ciao %s!", Name)
}

func (pr Portuguese) Greet(Name string) string {
	return fmt.Sprintf("Olá %s!", Name)
}
