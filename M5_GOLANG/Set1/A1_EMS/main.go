package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

const (
	HR = "HR"
	IT = "IT"
)

func main() {
	err := addEmployee(1, "Alice", 30, IT)
	if err != nil {
		fmt.Println(err)
	}
	err = addEmployee(2, "Bob", 25, HR)
	if err != nil {
		fmt.Println(err)
	}

	searchEmployeeByID(1)
	searchEmployeeByName("Bob")
	listEmployeesByDepartment(HR)
	countEmployeesInDepartment(IT)
}

func addEmployee(id int, name string, age int, department string) error {
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	fmt.Println("Employee added successfully!")
	return nil
}

func searchEmployeeByID(id int) {
	for _, emp := range employees {
		if emp.ID == id {
			fmt.Printf("Employee found: %+v\n", emp)
			return
		}
	}
	fmt.Println("Error: Employee not found.")
}

func searchEmployeeByName(name string) {
	for _, emp := range employees {
		if strings.EqualFold(emp.Name, name) {
			fmt.Printf("Employee found: %+v\n", emp)
			return
		}
	}
	fmt.Println("Error: Employee not found.")
}

func listEmployeesByDepartment(department string) {
	fmt.Printf("Employees in %s department:\n", department)
	for _, emp := range employees {
		if emp.Department == department {
			fmt.Printf("- %+v\n", emp)
		}
	}
}

func countEmployeesInDepartment(department string) {
	count := 0
	for _, emp := range employees {
		if emp.Department == department {
			count++
		}
	}
	fmt.Printf("Number of employees in %s department: %d\n", department, count)
}
