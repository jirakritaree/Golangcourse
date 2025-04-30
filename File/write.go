package main

import (
	"os"
)

func main() {
	data1 := []byte("Hello\n Jirakrit")
	err := os.WriteFile("File/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("File/employeeName")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data2 := []byte("Sira\n Manee")
	os.WriteFile("File/employeeName.txt", data2, 0644)
}
