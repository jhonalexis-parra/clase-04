package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Employee: %v \n", e)
}

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products = []Product{{
	ID:          1,
	Name:        "Arroz",
	Price:       20.0,
	Description: "Diana del Tolima",
	Category:    "Alimentos",
}, {
	ID:          2,
	Name:        "Verdura",
	Price:       16.0,
	Description: "Verdes",
	Category:    "Alimentos",
},
}

func (pm Product) Save() {
	Products = append(Products, pm)
}

func (pm Product) GetAll() {
	fmt.Println(Products)
}

func (pm Product) GetById(id int) {
	for i := 0; i < len(Products); i++ {
		if Products[i].ID == id {
			fmt.Println(Products[i])
			break
		}
	}
}

func main() {
	empleado1 := Employee{1, "Director", Person{1, "Jhon", "2013-01-01T00:00:00Z"}}
	empleado1.PrintEmployee()

	productTest := Product{
		ID:          3,
		Name:        "Aguacate",
		Price:       15.0,
		Description: "Amarillo",
		Category:    "Alimentos",
	}

	productTest.Save()
	productTest.GetAll()
	productTest.GetById(2)
}
