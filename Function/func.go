package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func plus(value1 int ,value2 int)int{
	return value1 + value2
}

func main() {
	hello()
	result := plus(5,100)
	fmt.Println("Result =", result)
}
