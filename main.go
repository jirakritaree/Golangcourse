package main

import "fmt"

var product = make(map[string]float64)

func main()  {
	fmt.Println("Product =", product)

	//add
	product["Macbook"] = 40000
	product["ipad"] = 20000
	product["iphone"] = 30000
	fmt.Println("Product =", product)

	//delete
	delete(product, "ipad")
	fmt.Println("Product =", product)

	//update
	product["Macbook"] = 39900
	product["ipad"] = 19900
	product["iphone"] = 24900
	fmt.Println("Product =", product)

	value := product["Macbook"]
	fmt.Println("Value =", value)
}