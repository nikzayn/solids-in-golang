package main

import "fmt"

type IPrinter interface {
	Print() string
}

type IScanner interface {
	Scan() string
}

type IMultiFunctionDevice interface {
	IPrinter
	IScanner
}

type Printer struct{}

func (p Printer) Print() string {
	return "Printing document..."
}

type ScannerImpl struct{}

func (p ScannerImpl) Scan() string {
	return "Scanning document..."
}

type MultiFunctionDeviceImpl struct{}

func (m MultiFunctionDeviceImpl) Print() string {
	return "Printing document..."
}

func (m MultiFunctionDeviceImpl) Scan() string {
	return "Scanning document..."
}

func main() {
	var m MultiFunctionDeviceImpl
	p := m.Print()
	s := m.Scan()
	fmt.Printf("Print function invoked: %v\n", p)
	fmt.Printf("Scan function invoked: %v\n", s)
}
