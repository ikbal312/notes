// package main

// import "fmt"

// // declaration
// var counters = make(map[string]int, 10)

// func main() {
// 	// declare using composite literal
// 	modelToMake := map[string]string{
// 		"prius":    "toyota",
// 		"chevelle": "chevy",
// 	}

// 	// accessing value
// 	// carMake := modelToMake["prius"]
// 	// fmt.Println(carMake)
// 	// accessing with unknown key
// 	// unknown := modelToMake["unknown"]
// 	// fmt.Println(unknown)

// 	// conditional check for key
// 	// if carMake, ok := modelToMake["unknown"]; ok {
// 	// 	fmt.Println("car model found", carMake)
// 	// } else {
// 	// 	fmt.Println("car not found with the key")
// 	// }

// 	// adding new value
// 	// modelToMake["outback"] = "subaru"
// 	// counters["pageHits"] = 10

// 	// extract all value

// 	for key, val := range modelToMake {
// 		fmt.Println(key, val)
// 	}

// }
