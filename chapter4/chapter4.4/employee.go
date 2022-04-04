package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerId int
	age       int
}

func EmployeeByID(id int) *Employee {
	return nil
}

func main() {
	employee := Employee{
		ID:  1,
		age: 18,
	}
	fmt.Println(employee.ID)
	fmt.Println(employee.age)
}
