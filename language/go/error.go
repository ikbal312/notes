// // go treat error  as data type
// // return nill means no error
// /*
//    package main

//    import (
//    	"errors"
//    	"fmt"
//    )

//    func Divide(num int, div int) (int, error) {
//    	if div == 0 {
//    		// we return the zero value of int (0) and an error.
//    		return 0, errors.New("cannot divide by 0")
//    	}
//    	return num / div, nil
//    }

//    func main() {
//    	divideBy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//    	for _, div := range divideBy {
//    		if res, err := Divide(100, div); err != nil {
//    			fmt.Printf("100 by %d error: %s\n", div, err)
//    		} else {
//    			fmt.Printf("100 divided by %d = %d\n", div, res)
//    		}
//    	}
//    }
// */

// /*
// // creating named error

// package main

// import (
// 	"errors"
// 	"log"
// 	"time"
// )

// // creating named error

// var (
// 	ErrNetwork = errors.New("network error")
// 	ErrInput   = errors.New("input error")
// )

// func main() {
// 	for {
// 		err := someFunc("data")
// 		if err == nil {
// 			// success so exit he loop
// 			break
// 		}

// 		if errors.Is(err, ErrNetwork) {
// 			log.Println("recoverable network error")
// 			time.Sleep(1 * time.Second)
// 			continue
// 		}
// 		log.Println("unrecoverable error")
// 		break // exit loop, as retrying useless
// 	}
// }

// func someFunc(data string) error {

// 	return ErrInput
// }
// */

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// )

// const (
// 	UnknownCode     = 0
// 	UnreachableCode = 1
// 	AuthFailureCode = 2
// )

// type ErrNetwork struct {
// 	Code int
// 	Msg  string
// }

// func (e ErrNetwork) Error() string {

// 	return fmt.Sprintf("network error(%d):%s", e.Code, e.Msg)
// }

// func userAuthCheck() error {

// 	// if auth filed for user
// 	return ErrNetwork{
// 		Code: AuthFailureCode,
// 		Msg:  "user unrecognized",
// 	}
// }
// func main() {

// 	for {

// 		err := userAuthCheck()
// 		var netError ErrNetwork
// 		if errors.As(err, &netError) {
// 			if netError.Code == AuthFailureCode {
// 				log.Println("unrecoverable auth failure:", err)
// 				break
// 			}
// 			log.Println("recoverable error:%s", netError)
// 		}
// 		log.Println("unrecoverable error: %s", err)
// 		break
// 	}
// }
