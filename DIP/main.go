package main

import "fmt"

type DataSource interface {
	GetData() string
}

// Database is a concrete implementation of DataSource
type Database struct{}

func (db Database) GetData() string {
	return "Data from the database"
}

// API is a high level module that depends on DataSource abstraction
type API struct {
	DataSource DataSource
}

func (api API) GetData() string {
	return api.DataSource.GetData()
}

func main() {
	db := Database{}
	api := API{DataSource: db}
	result := api.GetData()
	fmt.Println(result)
}
