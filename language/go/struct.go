// package main

// import "fmt"

// // struct represent a collection of variable
// type Record struct {
// 	Name string
// 	Age  int
// }

// // adding method to type

// func (r Record) Show() string {
// 	return fmt.Sprintf("%s,%d", r.Name, r.Age)
// }

// func ChangeName(r *Record) {
// 	r.Name = "Peter"
// 	fmt.Println("inside function", r.Name)
// }

// func (r *Record) IncAge() {
// 	r.Age++
// }

// func main() {

// 	r := &Record{Name: "ikbal", Age: 25}

// 	ChangeName(r)
// 	r.IncAge()
// 	fmt.Println(r)
// }
