package main

import "fmt"

func main() {

	fmt.Println("Hello World!!!")
	fmt.Println("For Test!!!")
	fmt.Println("For commit auto check!!!")
	fmt.Println("For SCM test.")
}

func Hello(name string) {
	fmt.Println("Hello ", name)
}

func Hello1(name string) string {
	return "Hello " + name
}
