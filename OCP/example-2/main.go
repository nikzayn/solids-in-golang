package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Legs() int {
	return 4
}

func (c Cat) DisplayLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type HexaCat struct {
	Cat
}

func (h HexaCat) Legs() int {
	return 6
}

func main() {
	var hex HexaCat
	fmt.Println(hex.Legs())
	hex.DisplayLegs()
}
