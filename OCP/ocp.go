package main

import "fmt"

type Greeter interface {
	Greet() string
}

type EnglishGreeter struct{}

type ItalianGreeter struct{}

type GermanGreeter struct{}

type HindiGreeter struct{}

type JapaneseGreeter struct{}

func (e EnglishGreeter) Greet() string {
	return "Hello"
}

func (i ItalianGreeter) Greet() string {
	return "Ciao"
}

func (g GermanGreeter) Greet() string {
	return "Hallo"
}

func (h HindiGreeter) Greet() string {
	return "नमस्ते"
}

func (j JapaneseGreeter) Greet() string {
	return "こんにちは"
}

func GreetPeople(greeter Greeter, names []string) {
	for _, name := range names {
		fmt.Println(greeter.Greet(), name)
	}
}

func main() {
	greets := []Greeter{
		&EnglishGreeter{},
		&ItalianGreeter{},
		&GermanGreeter{},
		&HindiGreeter{},
		&JapaneseGreeter{},
	}

	names := []string{"Alice", "Oslo", "Hitler", "Gandhi", "Emiko"}
	for _, greet := range greets {
		GreetPeople(greet, names)
	}
}
