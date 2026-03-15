package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// 变量
	name := "Go"
	version := "1.21"

	fmt.Printf("Welcome to %s %s\n", name, version)

	// 基本类型
	age := 25
	price := 19.99
	active := true

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Price: %.2f\n", price)
	fmt.Printf("Active: %t\n", active)
}
