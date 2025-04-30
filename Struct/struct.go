package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := [3]employee{}
	employeeList[0] = employee{
		employeeID:   "101",
		employeeName: "Prawet",
		phone:        "01234567890",
	}
	employeeList[1] = employee{
		employeeID:   "102",
		employeeName: "Prayad",
		phone:        "01234567891",
	}
	employeeList[2] = employee{
		employeeID:   "103",
		employeeName: "Prayut",
		phone:        "01234567892",
	}
	fmt.Println("Employee = ", employeeList)
}
