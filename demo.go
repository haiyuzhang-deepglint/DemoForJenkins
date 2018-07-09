package main

import (
	"fmt"
	"github.com/haiyuzhang-deepglint/Hello"
	"golang.org/x/sys/unix"
)

func main() {

	fmt.Println("Hello World!!!")
	fmt.Println("For Test!!!")
	fmt.Println("For commit auto check!!!")
	fmt.Println("For SCM test.")

	fmt.Println("Below message form Hello Package.")
	Hello.Hello("Jackie")
	fmt.Println(Hello.Hello1("Jenkins"))

	fmt.Println("Hello2 func :")
	Hello.Hello2("Jackie")
	Hello.Hello2("Jenkins")

	unix.ByteSliceFromString("Jenkins")
}
