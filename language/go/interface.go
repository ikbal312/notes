// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type Stringer interface {
// 	String() string
// }

// // person struct with method String
// type Person struct {
// 	First, Last string
// }

// // string method for Person Struct
// func (p Person) String() string {
// 	return fmt.Sprintf("%s,%s", p.Last, p.First)
// }

// // list of string
// type StrList []string

// func (s StrList) String() string {
// 	return strings.Join(s, ",")
// }

// // method to call struct function using interface

// func PrintStringer(s Stringer) {
// 	fmt.Println(s.String())
// }

// func main() {

// 	john := Person{First: "john", Last: "Doak"}

// 	var nameList Stringer = StrList{"David", "Sarah"}
// 	PrintStringer(john)
// 	PrintStringer(nameList)
// }
