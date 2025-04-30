package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {

	data, _ := json.Marshal(&employee{101, "Jirakrit Aree", "0992237646", "jirakritaree@gmail.com"})
	fmt.Println(string(data))
}
