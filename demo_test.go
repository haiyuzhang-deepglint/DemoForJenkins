package main

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T) {
	Hello("Jackie")
	Hello("Jenkins")
}

func TestHello1(t *testing.T) {
	result := Hello1("Jackie")
	if result!="Hello Jackie" {
		t.Fail()
	} else {
		fmt.Println("Hello1 return Hello Jackie.")
	}
}

func TestHello2(t *testing.T) {
	result := Hello1("Jenkins")
	if result!="Hello Jenkins" {
		t.Fail()
	} else {
		fmt.Println("Hello1 return Hello Jenkins.")
	}
}
