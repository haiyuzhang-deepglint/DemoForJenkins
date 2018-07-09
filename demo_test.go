package main

import (
	"github.com/haiyuzhang-deepglint/Hello"
	"testing"
	"fmt"
)

func TestHello(t *testing.T) {
	Hello.Hello("Jackie")
	Hello.Hello("Jenkins")
}

func TestHello1(t *testing.T) {
	result := Hello.Hello1("Jackie")
	if result!="Hello Jackie" {
		t.Fail()
	} else {
		fmt.Println("Hello1 return Hello Jackie.")
	}
}

func TestHello2(t *testing.T) {
	result := Hello.Hello1("Jenkins")
	if result!="Hello Jenkins" {
		t.Fail()
	} else {
		fmt.Println("Hello1 return Hello Jenkins.")
	}
}
