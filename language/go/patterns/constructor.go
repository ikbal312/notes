package main

import (
	"fmt"
	"log"
)

type Record struct {
	Name string
	Age  int
}

func NewRecord(name string, age int) (*Record, bool) {
	if name == "" {
		log.Println("name cannot be the empty string")
		return nil, false
	}
	if age <= 0 {
		log.Print("age should be 0<=age")
		return nil, false
	}
	return &Record{Name: name, Age: age}, true
}

func main() {
	if rec, ok := NewRecord("John Doak", 1); ok {
		fmt.Println(rec)
	} else {
		fmt.Println("error: user create failed")
	}

}
